package cloudtrail

import (
	"context"
	"fmt"

	"github.com/MauricioAntonioMartinez/aws/auth"
	"github.com/MauricioAntonioMartinez/aws/model"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *CloudTrailService) GetEvents(req *aws.GetEventsRequest, stream aws.CloudTrailService_GetEventsServer) error {

	return nil
}

func (c *CloudTrailService) CreateTrail(ctx context.Context, req *aws.CreateTrailRequest) (*aws.CreateTrailResponse, error) {

	us, err := auth.GetUserMetadata(ctx)
	if err != nil {
		c.logger.Err(err).Msg("User not found on context.")
		return nil, status.Errorf(codes.Unauthenticated, "cloudtrail: unauthenticated.")
	}

	arn, err := auth.NewArn(auth.CloudTrail, c.region, us.AccountId, fmt.Sprintf(":/trail/%s", req.GetName()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cloudtrail:unable to generate arn")
	}

	trail := model.Trail{
		Name:      req.Name,
		AccountId: us.AccountId,
		Arn:       arn.String(),
		TrailConfig: model.TrailConfig{
			Write: req.Config.Write,
			Read:  req.Config.Read,
		},
	}

	if err := repo.createTrail(&trail); err != nil {
		c.logger.Err(err)
		return nil, status.Errorf(codes.Internal, "cloudtrail: internal error trail not created.")
	}

	return &aws.CreateTrailResponse{
		Arn: arn.String(),
	}, nil
}

func (c *CloudTrailService) UpdateTrail(ctx context.Context, req *aws.CloudTrailConfig) (*aws.UpdateTrailResponse, error) {
	return nil, nil
}

func (c *CloudTrailService) DeleteTrail(ctx context.Context, req *aws.DeleteTrailRequest) (*aws.DeleteTrailResponse, error) {
	return nil, nil
}
