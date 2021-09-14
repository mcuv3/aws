package cloudtrail

import (
	"github.com/MauricioAntonioMartinez/aws/model"
	"gorm.io/gorm"
)

type cloudTrailRepo struct {
	db *gorm.DB
}

func (repo *cloudTrailRepo) createTrail(trail *model.Trail) error {
	tx := repo.db.Create(trail)
	return tx.Error
}

func (repo *cloudTrailRepo) createEvent(event *model.CloudTrailEvent) error {
	tx := repo.db.Create(event)
	return tx.Error
}

func (repo *cloudTrailRepo) updateTrail(event *model.Trail) error {
	tx := repo.db.Save(event)
	return tx.Error
}

func (repo *cloudTrailRepo) deleteTrail(trail model.Trail) error {
	tx := repo.db.Delete(trail)
	return tx.Error
}

// func (repo *cloudTrailRepo) getEvents(where model.CloudTrailEvent, events *[]model.CloudTrailEvent) error {
// 	tx := repo.db.Where(where).Find(events)
// 	return tx.Error
// }

func (repo *cloudTrailRepo) getEventsByQuery(query string, args ...interface{}) ([]model.CloudTrailEvent, error) {
	events := []model.CloudTrailEvent{}
	tx := repo.db.Where(query, args...).Find(&events)
	return events, tx.Error
}

func (repo *cloudTrailRepo) getTrails(where model.Trail, trails *[]model.Trail) error {
	tx := repo.db.Where(where).Find(trails)
	return tx.Error
}

func (repo *cloudTrailRepo) getTrail(where model.Trail, trails *model.Trail) error {
	tx := repo.db.Where(where).First(trails)
	return tx.Error
}
