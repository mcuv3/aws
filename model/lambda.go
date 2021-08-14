package model

import "gorm.io/gorm"


type Function struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);not null;unique_index"`
	Language string `gorm:"type:varchar(100);not null"`
	Runtime string `gorm:"type:varchar(50);not null"`
	Code string `gorm:"type:text;not null"`
	Image string `gorm:"type:varchar(50);not null"`
	AccountId string `gorm:"type:varchar(100);not null"`
} 


type Image struct {
	gorm.Model
	ImageName string `gorm:"type:varchar(100);not null;unique_index"`
	Versions []Version `gorm:"foreignKey:ImageID"`
}


type Version struct {
	gorm.Model
	ImageID uint 
	Tag string `gorm:"type:varchar(100);not null"`
}