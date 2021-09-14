package cloudtrail

import (
	"context"
	"fmt"
	"time"

	"github.com/MauricioAntonioMartinez/aws/auth"
	"github.com/MauricioAntonioMartinez/aws/model"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (c *CloudTrailService) GetEvents(req *aws.GetEventsRequest, stream aws.CloudTrailService_GetEventsServer) error {

	us, err := auth.GetUserMetadata(stream.Context())
	if err != nil {
		c.logger.Err(err).Msg("User not found on context.")
		return status.Errorf(codes.Unauthenticated, "cloudtrail: unauthenticated.")
	}

	now := time.Now()
	var secWithNoEvents int
	trails := []model.Trail{}

	if err := repo.getTrails(model.Trail{
		AccountId: us.AccountId,
	}, &trails); err != nil {
		c.logger.Err(err)
		return status.Errorf(codes.Internal, "cloudtrail: internal error trail not found.")
	}

	for {

		now = now.Add(time.Second)

		events := []model.CloudTrailEvent{}

		for _, trail := range trails {

			trailEvents, err := repo.getEventsByQuery("created_at  > ? AND trail_id = ?", now.Add(-time.Minute), trail.ID)
			if err != nil {
				c.logger.Err(err)
				continue
			}
			events = append(events, trailEvents...)
		}

		if len(events) == 0 {
			secWithNoEvents++
			if secWithNoEvents > 60 {
				break
			}
			continue
		}
		secWithNoEvents = 0

		for _, event := range events {
			stream.Send(&aws.CloudTrailEvent{
				Method:    event.Method,
				Timestamp: timestamppb.New(event.CreatedAt),
				TrailId:   int32(event.TrailID),
				Sid:       event.Sid,
				Region:    event.Region,
			})
		}

	}

	return nil
}

func (c *CloudTrailService) CreateTrail(ctx context.Context, req *aws.TrailRequest) (*aws.CreateTrailResponse, error) {

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

func (c *CloudTrailService) UpdateTrail(ctx context.Context, req *aws.TrailRequest) (*aws.UpdateTrailResponse, error) {
	us, err := auth.GetUserMetadata(ctx)
	if err != nil {
		c.logger.Err(err).Msg("User not found on context.")
		return nil, status.Errorf(codes.Unauthenticated, "cloudtrail: unauthenticated.")
	}

	trail := model.Trail{}

	if err := repo.getTrail(model.Trail{
		AccountId: us.AccountId,
		Name:      req.GetName(),
	}, &trail); err != nil {
		c.logger.Err(err)
		return nil, status.Errorf(codes.Internal, "cloudtrail: trail not found")
	}

	trail.TrailConfig = model.TrailConfig{
		Write: req.Config.Write,
		Read:  req.Config.Read,
	}

	if err := repo.updateTrail(&trail); err != nil {
		c.logger.Err(err)
		return nil, status.Errorf(codes.Internal, "cloudtrail: trail not updated")
	}

	return &aws.UpdateTrailResponse{
		Arn: trail.Arn,
	}, nil
}

func (c *CloudTrailService) DeleteTrail(ctx context.Context, req *aws.DeleteTrailRequest) (*aws.DeleteTrailResponse, error) {
	us, err := auth.GetUserMetadata(ctx)
	if err != nil {
		c.logger.Err(err).Msg("User not found on context.")
		return nil, status.Errorf(codes.Unauthenticated, "cloudtrail: unauthenticated.")
	}

	if err := repo.deleteTrail(model.Trail{
		AccountId: us.AccountId,
		Name:      req.GetName(),
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "cloudtrail: trail not deleted")
	}

	return &aws.DeleteTrailResponse{
		Ok:      true,
		Message: "Deleted successfully.",
	}, nil
}
