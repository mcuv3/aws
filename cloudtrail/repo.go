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

func (repo *cloudTrailRepo) getEvents(where model.CloudTrailEvent, events *[]model.CloudTrailEvent) error {
	tx := repo.db.Where(where).Find(events)
	return tx.Error
}

func (repo *cloudTrailRepo) getTrails(where model.Trail, trails *[]model.Trail) error {
	tx := repo.db.Where(where).Find(trails)
	return tx.Error
}
