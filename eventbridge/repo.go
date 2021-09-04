package eventbridge

import (
	"github.com/MauricioAntonioMartinez/aws/model"
	"gorm.io/gorm"
)

type EventBridgeRepo struct {
	db *gorm.DB
}

func (e *EventBridgeRepo) findServiceEvent(svc model.ServiceEvent) (model.ServiceEvent, error) {
	var svcEvent model.ServiceEvent
	if tx := e.db.Where(&svc).First(&svcEvent); tx.Error != nil {
		return svcEvent, tx.Error
	}
	return svcEvent, nil
}

func (e *EventBridgeRepo) findRule(rule *model.Rule) error {
	var r model.Rule
	if tx := e.db.Where(&rule).First(&r); tx.Error != nil {
		return tx.Error
	}
	(*rule) = r
	return nil
}

func (e *EventBridgeRepo) deleteRule(rule *model.Rule) error {
	tx := e.db.Delete(&rule)
	return tx.Error
}
