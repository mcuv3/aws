package eventbridge

import (
	"github.com/MauricioAntonioMartinez/aws/model"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
)

func (s *EventBridgeService) GetTargets(tgs []*aws.Target) []model.Target {
	targets := []model.Target{}

	for _, tg := range tgs {
		targets = append(targets, model.Target{
			Type: tg.TargetType.String(),
			Arn:  tg.GetTargetArn(),
		})
	}

	return targets
}

// Returns the associated gRPC prefix of a service based on its name
func (s *EventBridgeService) getServiceAlias(serviceName string) string {
	return ""
}
