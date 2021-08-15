package lambda

import (
	"context"
	"fmt"

	"github.com/MauricioAntonioMartinez/aws/model"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)


type Runtime struct {
	Image string
	Name string
	Extension string
	Activator string
	Version string
}

var (
	runtimes  = []Runtime{
		{Image: fmt.Sprintf("%s/node:%s",region,"14"), Name: "Nodejs 14.x", Version: "1",Extension: "js",Activator: "node" },
		{Image: fmt.Sprintf("%s/python:%s",region,"3.8"), Name: "Python 3.8", Version: "1",Extension: "py",Activator: "python"},
	}
)



func (l *LambdaServer) SeedLambdaServer(ctx context.Context,in *emptypb.Empty) (*aws.LambdaResponse, error) {

	rnts := []model.Runtime{}
	for _,runtime := range runtimes {
		rnts = append(rnts, model.Runtime{Image: runtime.Image, Name: runtime.Name, Version: runtime.Version})
	}

	tx := l.db.CreateInBatches(rnts,len(runtimes))

	if tx.Error !=nil {
		return nil,grpc.Errorf(codes.Internal,"Couldn't seed the runtimes")
	}

	return &aws.LambdaResponse{Message: "Sucessfully seeded the runtimes",Ok: true},nil
}
