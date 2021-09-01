package eventbridge

import (
	"context"

	aws "github.com/MauricioAntonioMartinez/aws/proto"
)

func (s *EventBridgeService) CreateRule(ctx context.Context, req *aws.CreateRuleRequest) (*aws.CreateRuleResponse, error) {
	return nil, nil
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
