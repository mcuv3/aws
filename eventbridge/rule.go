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

func (s *EventBridgeService) CreateRule(ctx context.Context, req *aws.CreateRuleRequest) (*aws.CreateRuleResponse, error) {
	us, err := auth.GetUserMetadata(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "401:eventbridge")
	}

	arn, err := auth.NewArn(auth.Service(s.Name), s.region, us.AccountId, fmt.Sprintf("/rule/%s", req.GetName()))
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
	err = repo.createRule(&rule)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "500:could-not-save-rule")
	}

	return &aws.CreateRuleResponse{
		RuleArn: rule.Arn,
	}, nil
}

func (s *EventBridgeService) UpdateRule(ctx context.Context, req *aws.UpdateRuleRequest) (*aws.EventBridgeResponse, error) {
	_, err := s.getRuleForUser(ctx, req.GetName())
	if err != nil {
		return nil, err
	}

	cron, svcEvent, err := s.getRuleEvents(rulePayload{scheduleEvent: req.GetScheduleEvent(), eventPattern: req.GetEventPattern()})
	if err != nil {
		return nil, err
	}

	fmt.Println("cron ", cron)

	newRule := model.Rule{
		Description:    req.GetDescription(),
		ServiceEventID: svcEvent,
		EventPattern:   cron,
		Active:         true,
	}

	updates := []string{"description", "service_event_id", "event_pattern", "active"}
	if err = repo.updateRule(&newRule, updates...); err != nil {
		return nil, status.Errorf(codes.Internal, "500:could-not-update-rule")
	}

	return &aws.EventBridgeResponse{
		Message: "Successfully updated",
		Status:  "200",
	}, nil
}

// Deletes a rule from the datbase
func (s *EventBridgeService) DeleteRule(ctx context.Context, req *aws.DeleteRuleRequest) (*aws.EventBridgeResponse, error) {
	rule, err := s.getRuleForUser(ctx, req.GetName())
	if err != nil {
		return nil, err
	}

	if err := repo.deleteRule(rule); err != nil {
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

	if err := repo.updateRule(rule); err != nil {
		return nil, status.Errorf(codes.Internal, "500:rule-toggle-active")
	}

	return &aws.EventBridgeResponse{
		Message: "Successfully updated status.",
		Status:  "200",
	}, nil
}

func (s *EventBridgeService) UpdateTarget(ctx context.Context, req *aws.UpdateTargetRequest) (*aws.EventBridgeResponse, error) {
	return nil, nil
}

func (s *EventBridgeService) getRuleForUser(ctx context.Context, rulename string) (*model.Rule, error) {
	us, err := auth.GetUserMetadata(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "401:eventbridge-unauthenticated")
	}

	rule := model.Rule{
		Name:      rulename,
		AccountId: us.AccountId,
	}
	if err = repo.findRule(&rule); err != nil {
		return nil, status.Errorf(codes.NotFound, "404:rule")
	}

	return &rule, nil
}

type rulePayload struct {
	scheduleEvent *aws.ScheduleEvent
	eventPattern  *aws.EventPattern
}

func (s *EventBridgeService) getRuleEvents(req rulePayload) (*string, *uint, error) {
	cron := req.scheduleEvent.GetCron()
	event := req.eventPattern
	svcMethod := event.GetEventType()
	svc := event.GetService().String()

	isServiceEvent := (svcMethod != "" && svc != "")
	isCronEvent := (cron != "")
	if (isCronEvent && isServiceEvent) || (!isCronEvent && !isServiceEvent) {
		return nil, nil, status.Errorf(codes.InvalidArgument, "400:only-one-event")
	}

	svcEvent := model.ServiceEvent{
		Name:   strings.ToLower(svc),
		Method: svcMethod,
	}
	var cronEvent *string

	if isCronEvent {
		cronEvent = &cron
	}

	if isServiceEvent {
		if err := repo.findServiceEvent(&svcEvent); err != nil {
			return nil, nil, status.Errorf(codes.InvalidArgument, "Unknown source or type")
		}
	}

	return cronEvent, &svcEvent.ID, nil

}
