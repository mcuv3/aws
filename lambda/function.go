package lambda

import (
	"context"

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

	
	function := model.Function{
		Name: req.GetName(),
		Language: req.GetLanguage(),
		Runtime: req.GetRuntime(),
		Code: req.GetCode(),
		Image: image,
		AccountId: us.AccountId,
	}
	
	tx :=  l.db.Create(&function)
	
	if tx.Error != nil {
		return nil,grpc.Errorf(codes.Internal,"Error creating function.")
	}
	
	l.docker.BuildImage(docker.BuildImageOptions{
		ImageName: image,
		Code: req.GetCode(),
		Language: lang,
	})

	return &aws.LambdaResponse{Message: "Successfully function created",Ok: true},nil
}

func (l *LambdaServer) TestFunction(ctx context.Context, req *aws.TestFunctionResquest) (*aws.LambdaResponse,error){

	_ = docker.DockerRuntime{}


	return nil,nil
 }