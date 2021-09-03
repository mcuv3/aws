package eventbridge

import (
	"context"
	"fmt"
	"strings"

	"github.com/MauricioAntonioMartinez/aws/auth"
	"github.com/MauricioAntonioMartinez/aws/model"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TODO: test is this rpc call work.

func (s *EventBridgeService) CreateRule(ctx context.Context, req *aws.CreateRuleRequest) (*aws.CreateRuleResponse, error) {

	us, err := s.auth.GetUserMetadata(ctx)

	s.logger.Debug().Msgf("User metadata: %+v", us)
	s.logger.Debug().Msgf("Err: %v", err)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "401:eventbridge")
	}

	arn, err := auth.NewArn(s.Name, s.region, us.AccountId, fmt.Sprintf("/rule/%s", req.GetName()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "500:ARN")
	}

	cron, svcEvent, err := s.getRuleEvents(rulePayload{scheduleEvent: req.GetScheduleEvent(), eventPattern: req.GetEventPattern()})
	if err != nil {
		return nil, err
	}

	rule := model.Rule{
		Name:           req.GetName(),
		Description:    req.GetDescription(),
		Arn:            arn.String(),
		AccountId:      us.AccountId,
		ServiceEventID: svcEvent,
		EventPattern:   cron,
		Targets:        s.GetTargets(req.GetTargets()),
	}

	tx := s.db.Save(&rule)

	if tx.Error != nil {
		return nil, status.Errorf(codes.Internal, "500:could-not-save-rule")
	}

	return &aws.CreateRuleResponse{
		RuleArn: rule.Arn,
	}, nil
}

func (s *EventBridgeService) UpdateRule(ctx context.Context, req *aws.CreateRuleRequest) (*aws.EventBridgeResponse, error) {
	rule, err := s.getRuleForUser(ctx, req.GetName())
	if err != nil {
		return nil, err
	}

	cron, svcEvent, err := s.getRuleEvents(rulePayload{scheduleEvent: req.GetScheduleEvent(), eventPattern: req.GetEventPattern()})
	if err != nil {
		return nil, err
	}

	newRule := model.Rule{
		Description:    rule.Description,
		ServiceEventID: svcEvent,
		EventPattern:   cron,
		Targets:        s.GetTargets(req.GetTargets()), // this will the delete the old ones? if not delete all of them manually.
	}

	tx := s.db.Save(&newRule)

	if tx.Error != nil {
		return nil, status.Errorf(codes.Internal, "500:could-not-update-rule")
	}

	return nil, nil
}

// Deletes a rule from the datbase
func (s *EventBridgeService) DeleteRule(ctx context.Context, req *aws.DeleteRuleRequest) (*aws.EventBridgeResponse, error) {
	rule, err := s.getRuleForUser(ctx, req.GetName())
	if err != nil {
		return nil, err
	}

	if tx := s.db.Delete(&rule); tx.Error != nil {
		return nil, status.Errorf(codes.Internal, "500:rule-not-deleted")
	}

	return &aws.EventBridgeResponse{Message: "Successfully deleted", Status: "200"}, nil
}

// Changes the state of a rule
func (s *EventBridgeService) ChangeRuleState(ctx context.Context, req *aws.ChangeRuleStateRequest) (*aws.EventBridgeResponse, error) {
	rule, err := s.getRuleForUser(ctx, req.GetName())
	if err != nil {
		return nil, err
	}
	rule.Active = !rule.Active

	if tx := s.db.Save(&rule); tx.Error != nil {
		return nil, status.Errorf(codes.Internal, "500:rule-toggle-active")
	}

	return nil, nil
}

func (s *EventBridgeService) getRuleForUser(ctx context.Context, rulename string) (*model.Rule, error) {
	us, err := s.auth.GetUserMetadata(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "401:eventbridge-unauthenticated")
	}

	rule := model.Rule{}
	if tx := s.db.Where("name = ?", rulename).First(&rule); tx.Error != nil || strings.Contains(rule.Arn, us.AccountId) {
		return nil, status.Errorf(codes.NotFound, "404:rule")
	}
	return &rule, nil
}

type rulePayload struct {
	scheduleEvent *aws.ScheduleEvent
	eventPattern  *aws.EventPattern
}

func (s *EventBridgeService) getRuleEvents(req rulePayload) (string, *uint, error) {
	cron := req.scheduleEvent.GetCron()
	event := req.eventPattern
	eType := event.GetEventType()
	eSource := event.GetService().String()

	isServiceEvent := (eType != "" && eSource != "")
	isCronEvent := (cron != "")
	if (isCronEvent && isServiceEvent) || (!isCronEvent && !isServiceEvent) {
		return "", nil, status.Errorf(codes.InvalidArgument, "400:only-one-event")
	}

	var svcEvent *uint

	if isServiceEvent {
		svcE := model.ServiceEvent{}
		if err := s.db.Where("path = ?", fmt.Sprintf("%s/%s", eSource, eType)).First(&svcE); err != nil {
			return "", nil, status.Errorf(codes.InvalidArgument, "Unknown source or type")
		}
		svcEvent = &svcE.ID
	}

	return cron, svcEvent, nil

}
