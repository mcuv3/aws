package lambda

import (
	"context"
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/MauricioAntonioMartinez/aws/auth"
	"github.com/MauricioAntonioMartinez/aws/container"
	"github.com/MauricioAntonioMartinez/aws/model"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (l *LambdaService) CreateFunction(ctx context.Context, req *aws.CreateFunctionRequest) (*aws.LambdaResponse, error) {
	us, err := auth.GetUserMetadata(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Unable to authenticate in lambda.")
	}

	image := strings.ToLower("mcuve" + "/" + l.region + "-" + us.AccountId + "-" + req.GetName())
	runtime := model.Runtime{}

	tx := l.db.Where("name = ?", req.GetRuntime().String()).First(&runtime)
	if tx.Error != nil {
		return nil, status.Errorf(codes.Internal, "Unable to find runtime")
	}

	ext := path.Ext(req.GetHandler())
	p := strings.Replace(req.GetHandler(), ext, "", 1)

	arn, err := auth.NewArn(auth.Lambda, l.region, us.AccountId, req.GetName())

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to generare arn")
	}

	function := model.Function{
		Name:      req.GetName(),
		RuntimeID: runtime.ID,
		Code:      req.GetCode(),
		Image:     image,
		Handler:   req.GetHandler(),
		AccountId: us.AccountId,
		Path:      p,
		Arn:       arn.String(),
	}

	tx = l.db.Create(&function)
	if tx.Error != nil {
		return nil, status.Errorf(codes.Internal, "Error creating function.")
	}

	dockerFile := fmt.Sprintf(` 
	FROM %s
	RUN echo "%s" > %s.%s 

	RUN ls
	CMD ["%s", "runtime.%s"]
	`, runtime.Image, req.GetCode(), p, runtime.Extension, runtime.Activator, runtime.Extension)

	dispatcher.BuildImage(container.BuildImageOptions{
		ImageName:  image,
		BaseImage:  runtime.Image,
		Code:       req.GetCode(),
		DockerFile: dockerFile,
	})

	return &aws.LambdaResponse{Message: "Successfully function created", Ok: true}, nil
}

func (l *LambdaService) TestFunction(ctx context.Context, req *aws.TestFunctionResquest) (*aws.LambdaResponse, error) {

	us, err := auth.GetUserMetadata(ctx)

	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Unable to authenticate from lambda")
	}

	image := strings.ToLower("mcuve" + "/" + l.region + "-" + us.AccountId + "-" + req.GetFunctionName())

	res := model.Function{}

	tx := l.db.Where("image = ?", image).First(&res)

	if tx.Error != nil {
		return nil, status.Errorf(codes.NotFound, "Function not found")
	}

	if err := execution.Run(LambdaExecutionConfig{
		ID:      res.Arn,
		image:   res.Image,
		ram:     int64(res.Memory),
		data:    req.GetEventData(),
		handler: res.Handler,
	}); err != nil {
		l.logger.Err(err)
		return nil, status.Errorf(codes.Internal, "Unable to start the execution.")
	}

	return &aws.LambdaResponse{
		Message: fmt.Sprintf("Ran  at %s", time.Now()),
		Ok:      true,
	}, nil
}

func (l *LambdaService) InvokeFunction(ctx context.Context, req *aws.InvoqueFunctionRequest) (*aws.LambdaResponse, error) {

	res := model.Function{}

	tx := l.db.Where("arn = ?", req.GetArn()).First(&res)

	if tx.Error != nil {
		return nil, status.Errorf(codes.NotFound, "Function not found")
	}

	if err := execution.Run(LambdaExecutionConfig{
		ID:      res.Arn,
		image:   res.Image,
		ram:     int64(res.Memory),
		data:    req.GetEventData(),
		handler: res.Handler,
	}); err != nil {
		l.logger.Err(err)
		return nil, status.Errorf(codes.Internal, "Unable to start the execution.")
	}

	return &aws.LambdaResponse{
		Message: fmt.Sprintf("Ran at %s", time.Now()),
		Ok:      true,
	}, nil
}

func (l *LambdaService) ReceiveEvents(req *aws.ReceiveEventRequest, stream aws.LambdaService_ReceiveEventsServer) error {

	execution := execution.getExecutionByHash(req.GetHash())

	if execution == nil {
		return status.Errorf(codes.Aborted, "Invalid or unknown hash")
	}

	fn := model.Function{}

	tx := l.db.Where("arn = ?", execution.ID).First(&fn)

	if tx.Error != nil {
		return status.Errorf(codes.NotFound, "Function not found")
	}

broker:
	for {
		fmt.Println("Starting sending messages to lambda runtimes")
		select {
		case event := <-execution.Events:
			stream.Send(&aws.EventResponse{Message: event})
			execution.mt.Lock()
			execution.CurrentExecutions--
			execution.mt.Unlock()
		case <-time.After(time.Second * 30):
			// remove from the list of running executions
			l.logger.Info().Str("Expired receive", execution.ValHash)
			break broker
		}

	}

	fmt.Println("Connection closed from server hash ", execution.ValHash)

	return nil
}

func (l *LambdaService) UpdateLambda(ctx context.Context, req *aws.UpdateLambdaRequest) (*aws.LambdaResponse, error) {

	lambda := model.Function{}

	tx := l.db.Where("arn = ?", req.GetArn()).First(&lambda)

	if tx.Error != nil {
		return nil, status.Errorf(codes.NotFound, "Invalid function")
	}

	if req.GetCode() != "" {
		lambda.Code = req.GetCode()
	}
	if req.GetRam() != 0 {
		lambda.Memory = int(req.GetRam())
	}

	if req.GetHandler() != "" {
		lambda.Handler = req.GetHandler()
	}

	tx = l.db.Save(&lambda)

	if tx.Error != nil {
		return nil, status.Errorf(codes.Internal, "Unable to update the lambda function")
	}

	return &aws.LambdaResponse{
		Message: "Successfully updated lambda function",
		Ok:      true,
	}, nil
}

func (l *LambdaService) DeleteLambda(ctx context.Context, req *aws.DeleteLambdaRequest) (*aws.LambdaResponse, error) {
	lambda := model.Function{}

	tx := l.db.Where("arn = ?", req.GetArn()).First(&lambda)

	if tx.Error != nil {
		return nil, status.Errorf(codes.NotFound, "Invalid function")
	}

	tx = l.db.Delete(&lambda)

	if tx.Error != nil {
		return nil, status.Errorf(codes.Internal, "Unable to delete the function please try later")
	}

	return &aws.LambdaResponse{
		Message: "Successfully deleted lambda function",
		Ok:      true,
	}, nil
}
