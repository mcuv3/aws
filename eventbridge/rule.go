package eventbridge

import (
	"context"
	"fmt"

	"github.com/MauricioAntonioMartinez/aws/auth"
	"github.com/MauricioAntonioMartinez/aws/model"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func (s *EventBridgeService) CreateRule(ctx context.Context, req *aws.CreateRuleRequest) (*aws.CreateRuleResponse, error) {

	us, err := s.auth.GetUserMetadata(ctx)

	s.logger.Debug().Msgf("User metadata: %+v", us)
	s.logger.Debug().Msgf("Err: %v", err)
	if err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "401:eventbridge")
	}

	arn, err := auth.NewArn(s.Name, s.region, us.AccountId, fmt.Sprintf("/rule/%s", req.GetName()))
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "500:ARN")
	}

	cron := req.GetScheduleEvent().GetCron()
	event := req.GetEventPattern()
	eType := event.GetEventType()
	eSource := event.GetService().String()

	isServiceEvent := (eType != "" && eSource != "")
	isCronEvent := (cron != "")
	if (isCronEvent && isServiceEvent) || (!isCronEvent && !isServiceEvent) {
		return nil, grpc.Errorf(codes.InvalidArgument, "400:only-one-event")
	}

	var svcEvent *uint

	if isServiceEvent {
		svcE := model.ServiceEvent{}
		if err := s.db.Where("path = ?", fmt.Sprintf("%s/%s", eSource, eType)).First(&svcE); err != nil {
			return nil, grpc.Errorf(codes.InvalidArgument, "Unknown source or type")
		}
		svcEvent = &svcE.ID
	}

	rule := model.Rule{
		Name:           req.GetName(),
		Description:    req.GetDescription(),
		Arn:            arn.String(),
		ServiceEventID: svcEvent,
		EventPattern:   cron,
		Targets:        s.GetTargets(req.GetTargets()),
	}

	tx := s.db.Save(&rule)

	if tx.Error != nil {
		return nil, grpc.Errorf(codes.Internal, "500:could-not-save-rule")
	}

	return &aws.CreateRuleResponse{
		RuleArn: rule.Arn,
	}, nil
}

func (s *EventBridgeService) UpdateRule(ctx context.Context, req *aws.CreateRuleRequest) (*aws.EventBridgeResponse, error) {
	return nil, nil
}

func (s *EventBridgeService) DeleteRule(ctx context.Context, req *aws.DeleteRuleRequest) (*aws.EventBridgeResponse, error) {
	return nil, nil
}

func (s *EventBridgeService) ChangeRuleState(ctx context.Context, req *aws.ChangeRuleStateRequest) (*aws.EventBridgeResponse, error) {
	return nil, nil
}
