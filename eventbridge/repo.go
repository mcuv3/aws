package eventbridge

import (
	"database/sql"

	"github.com/MauricioAntonioMartinez/aws/model"
	"gorm.io/gorm"
)

type EventBridgeRepo struct {
	db *gorm.DB
}

func (e *EventBridgeRepo) findServiceEvent(svc *model.ServiceEvent) error {
	if tx := e.db.Where(&svc).First(&svc); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (e *EventBridgeRepo) findRule(rule *model.Rule) error {
	var r model.Rule
	if tx := e.db.Where(&rule).First(&r); tx.Error != nil {
		return tx.Error
	}
	(*rule) = r
	return nil
}

func (e *EventBridgeRepo) findRules(where model.Rule, rules *[]model.Rule) error {
	if tx := e.db.Where(where).Find(&rules); tx.Error != nil {
		return tx.Error
	}
	return nil
}

// func (e *EventBridgeRepo) findRulesQuery(query string, rules *[]model.Rule, args ...interface{}) error {
// 	if tx := e.db.Where(query, args...).Find(&rules); tx.Error != nil {
// 		return tx.Error
// 	}
// 	return nil
// }

func (e *EventBridgeRepo) findRulesWithTargets(accountId string, serviceEventID *uint) (*sql.Rows, error) {
	// if tx := e.db.Joins("rules", e.db.Where(&where)).Find(targets); tx.Error != nil {
	// 	return tx.Error
	// }
	// return nil
	rows, err := e.db.Table("targets").
		Select("targets.rule_id,targets.type,targets.arn,rules.arn,rules.event_pattern").
		Joins("inner join rules on rules.id = targets.rule_id").
		Where("rules.account_id = ? AND rules.service_event_id = ?", accountId, *serviceEventID).
		Rows()
	return rows, err
}

func (e *EventBridgeRepo) deleteRule(rule *model.Rule) error {
	tx := e.db.Delete(&rule)
	return tx.Error
}

func (e *EventBridgeRepo) createServiceEvent(svc *model.ServiceEvent) error {
	tx := e.db.Create(svc)
	return tx.Error
}

func (e *EventBridgeRepo) updateRule(rule *model.Rule, updates ...string) error {

	var err error

	if len(updates) == 0 {
		tx := e.db.Save(&rule)
		err = tx.Error
	} else {
		tx := e.db.Model(&rule).Select(updates).Updates(rule)
		err = tx.Error
	}

	return err
}

func (e *EventBridgeRepo) createRule(rule *model.Rule) error {
	tx := e.db.Create(rule)
	return tx.Error
}
