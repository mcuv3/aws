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
