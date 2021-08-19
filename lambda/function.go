package lambda

import (
	"context"
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/MauricioAntonioMartinez/aws/auth"
	"github.com/MauricioAntonioMartinez/aws/docker"
	"github.com/MauricioAntonioMartinez/aws/model"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)



func (l *LambdaServer) CreateFunction(ctx context.Context,req *aws.CreateFunctionRequest) (*aws.LambdaResponse,error) {
	us,err := l.auth.GetUserMetadata(ctx)
	if err !=nil {
		return nil,grpc.Errorf(codes.Unauthenticated,"Unable to authenticate in lambda.")
	}

	image := strings.ToLower("mcuve" +"/"+l.region + "-" + us.AccountId + "-" + req.GetName())
	runtime := model.Runtime{}

	tx := l.db.Where("name = ?",req.GetRuntime().String()).First(&runtime)
	if tx.Error !=nil {
		return nil,grpc.Errorf(codes.Internal,"Unable to find runtime")
	} 
	path.Base(req.GetHandler())
	ext := path.Ext(req.GetHandler())
	p := strings.Replace(req.GetHandler(),ext,"",1)

	arn, err := auth.NewArn(auth.Lambda,l.region,us.AccountId,req.GetName())

	if err != nil {
		return nil,grpc.Errorf(codes.Internal,"Unable to generare arn")
	}
 
	function := model.Function{
		Name: req.GetName(),
		RuntimeID: runtime.ID, 
		Code: req.GetCode(),
		Image: image, 
		Handler: req.GetHandler(),
		AccountId: us.AccountId,
		Path: p,
		Arn: arn.String(),
	}
	
	tx =  l.db.Create(&function)
	if tx.Error != nil {
		return nil,grpc.Errorf(codes.Internal,"Error creating function.")
	}


	dockerFile := fmt.Sprintf(`
	FROM %s
	RUN echo "%s" > %s.%s 
	CMD ["%s", "runtime.%s"]
	`,runtime.Image,req.GetCode(),p,runtime.Extension,runtime.Activator,runtime.Extension)

	
	l.docker.BuildImage(docker.BuildImageOptions{ 
		ImageName: image,
		BaseImage: runtime.Image,
		Code: req.GetCode(),
		DockerFile: dockerFile,
	})

	return &aws.LambdaResponse{Message: "Successfully function created",Ok: true},nil
}

func (l *LambdaServer) TestFunction(ctx context.Context, req *aws.TestFunctionResquest) (*aws.LambdaResponse,error){

	us, err := l.auth.GetUserMetadata(ctx)

	if err != nil {
		return nil,grpc.Errorf(codes.Unauthenticated,"Unable to authenticate from lambda")
	}


	image := strings.ToLower("mcuve" +"/"+l.region + "-" + us.AccountId + "-" + req.GetFunctionName())
	
	res := model.Function{}

	tx := l.db.Where("image = ?",image).First(&res)

	if  tx.Error !=nil {
		return nil,grpc.Errorf(codes.NotFound,"Function not found")
	}
 
	env := map[string]interface{}{
		"EVENT_DATA": req.GetEventData(),
		"HANDLER" : res.Handler,
	}


	l.docker.RunContainer(docker.RunContainerOptions{
		Image: res.Image,
		Ram: int64(res.Memory),
		Environment: env,
		Name: fmt.Sprintf("%s-%s-%d",us.AccountId,res.Name,time.Now().Unix()),
	})


	return &aws.LambdaResponse{
		Message: fmt.Sprintf("Ran  at %s",time.Now()),
		Ok:true,
	},nil
}


func (l *LambdaServer) InvoqueFunction(ctx context.Context, req *aws.InvoqueFunctionRequest) (*aws.LambdaResponse,error) {

	res := model.Function{}

	tx := l.db.Where("arn = ?",req.GetArn()).First(&res)

	if  tx.Error !=nil {
		return nil,grpc.Errorf(codes.NotFound,"Function not found")
	}

 
	env := map[string]interface{}{
		"EVENT_DATA": req.GetEventData(),
		"HANDLER" : res.Handler,
	}


	freeExecution := l.getAvaibleExecution(res.Arn)

	if freeExecution == nil {
		l.docker.RunContainer(docker.RunContainerOptions{
			Image: res.Image,
			Ram: int64(res.Memory),
			Environment: env,
		})
		l.newLambdaExecution(res.Arn)
	} else {
		go func(){
			freeExecution.Events <- req.GetEventData()
		}()
	}



	return &aws.LambdaResponse{
		Message: fmt.Sprintf("Ran at %s",time.Now()),
		Ok:true,
	},nil
}


func (l *LambdaServer) ReceiveEvents( req *aws.ReceiveEventRequest,stream aws.LambdaService_ReceiveEventsServer) error  {
	

	execution := l.LambdaExecutionManager.getExection(req.GetHash())

	if execution == nil {
		return grpc.Errorf(codes.Aborted,"Invalid or unknown token")
	}

	fn := model.Function{}

	tx := l.db.Where("arn = ?",execution.FuncArn).First(&fn)

	if tx.Error !=nil {
		return grpc.Errorf(codes.NotFound,"Function not found")
	}

	broker :for {
		select {
			case event := <- execution.Events:
				stream.Send(&aws.EventResponse{Message: event})
				execution.mt.Lock()
				execution.CurrentExecutions-- 
				execution.mt.Unlock()
			case  <- time.After(time.Second * 60):
				l.logger.Info().Str("Expired receive",execution.ValHash)
				break broker
		}

	}


	return nil
}