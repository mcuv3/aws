package sqs

import (
	"context"
	"errors"
	"fmt"

	"github.com/MauricioAntonioMartinez/aws/auth"
	"github.com/MauricioAntonioMartinez/aws/model"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *SQSService) CreateQueue(ctx context.Context, req *aws.CreateQueueRequest) (*aws.CreateQueueResponse, error) {
	queueName := req.GetName()

	us, err := auth.GetUserMetadata(ctx)

	if err != nil {
		return nil, s.Error(err, codes.Unauthenticated, "Unable to authenticate")
	}

	if res := s.db.Where("account_id = ? AND name =  ?", us.AccountId, queueName).First(&model.Queue{}); res.Error == nil {
		return nil, status.Errorf(codes.NotFound, "Something went wrong %s", queueName)
	}

	conf := req.GetConf()

	if string(conf.DeliveryDelayTime) == "" || string(conf.MessageRetentionTime) == "" || string(conf.VisibilityTime) == "" {
		return nil, s.Error(errors.New("Invalid configuration"), codes.InvalidArgument, "Invalid configuration please check.")
	}

	arn, err := auth.NewArn(auth.SQS, auth.US_EAST_1, us.AccountId, fmt.Sprintf("/queue/%s", queueName))

	if err != nil {
		return nil, s.Error(err, codes.InvalidArgument, "Invalid ARN")
	}

	queue := model.Queue{
		Name:      queueName,
		AccountId: us.AccountId,
		Arn:       arn.String(),
		Configuration: model.ConfigurationQueue{
			MessageRetentionTime: conf.MessageRetentionTime.String(),
			VisibilityTimeout:    int(conf.VisibilityTimeout),
			VisibilityTime:       conf.VisibilityTime.String(),
			MessageRetention:     int(conf.MessageRetention),
			DeliveryDelay:        int(conf.DeliveryDelayTime),
			DeliveryDelayTime:    conf.DeliveryDelayTime.String(),
		},
	}

	res := s.db.Create(&queue)

	if res.Error != nil {
		return nil, s.Error(res.Error, codes.Aborted, "Couldn't create queue.")
	}

	return &aws.CreateQueueResponse{Queue: &aws.Queue{
		Name: queueName,
		Arn:  arn.String(),
	}}, nil
}

func (s *SQSService) DeleteQueue(ctx context.Context, req *aws.DeleteQueueRequest) (*aws.DeleteResponse, error) {

	return nil, nil
}
