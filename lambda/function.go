package lambda

import (
	"context"
	"fmt"

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

	function := model.Function{
		Name: req.GetName(),
		Language: req.GetLanguage(),
		Runtime: req.GetRuntime(),
		Code: req.GetCode(),
		AccountId: us.AccountId,
	}

	tx :=  l.db.Create(&function)

	if tx.Error != nil {
		return nil,grpc.Errorf(codes.Internal,"Error creating function.")
	}

	fmt.Println(us)

	return &aws.LambdaResponse{Message: "Successfully function created",Ok: true},nil
}

func (l *LambdaServer) TestFunction(ctx context.Context, req *aws.TestFunctionResquest) (*aws.LambdaResponse,error){
	return nil,nil
 }