package lambda

import (
	"context"
	"fmt"
	"strings"

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

	image := l.region + "/" + us.AccountId + "-" + req.GetName()

	lang,ok := model.AllowedLanguages[req.GetLanguage()]

	if !ok {
		return nil,grpc.Errorf(codes.InvalidArgument,"Language not available")
	}

	runtime := model.Runtime{}

	tx := l.db.Where("runtime_name = ?",req.GetRuntime()).First(&runtime)

	if tx.Error !=nil {
		return nil,grpc.Errorf(codes.Internal,"Unable to find runtime")
	}


	function := model.Function{
		Name: req.GetName(),
		Language: req.GetLanguage(),
		RuntimeID: runtime.ID,
		Code: req.GetCode(),
		Image: image,
		AccountId: us.AccountId,
	}
	
	tx =  l.db.Create(&function)
	
	if tx.Error != nil {
		return nil,grpc.Errorf(codes.Internal,"Error creating function.")
	}

	dockerFile := fmt.Sprintf(`
	FROM %s
	RUN echo "%s" > code.%s
	CMD ["%s", "code.%s"]
	`,runtime.Image,req.GetCode(),runtime.Extension,runtime.Activator,runtime.Extension)
	
	l.docker.BuildImage(docker.BuildImageOptions{
		ImageName: image,
		BaseImage: runtime.Image,
		Code: req.GetCode(),
		DockerFile: dockerFile,
		Language: lang,
	})

	return &aws.LambdaResponse{Message: "Successfully function created",Ok: true},nil
}

func (l *LambdaServer) TestFunction(ctx context.Context, req *aws.TestFunctionResquest) (*aws.LambdaResponse,error){

	us, err := l.auth.GetUserMetadata(ctx)

	if err != nil {
		return nil,grpc.Errorf(codes.Unauthenticated,"Unable to authenticate from lambda")
	}

	functionName := strings.ToLower(l.region+"/"+us.AccountId + "-" + req.GetFunctionName())

	fn := model.Function{}

	tx := l.db.Where("name = ?",functionName).First(&fn)

	if  tx.Error !=nil {
		return nil,grpc.Errorf(codes.NotFound,"Function not found")
	}

	l.docker.RunContainer(docker.RunContainerOptions{
		Image: fn.Image,

	})



	return nil,nil
}


func (l *LambdaServer) InvoqueFunction(ctx context.Context, req *aws.InvoqueFunctionRequest) (*aws.LambdaResponse,error) {
	return nil,nil
}