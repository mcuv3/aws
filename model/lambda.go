package model

import "gorm.io/gorm"


type Function struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);not null;unique_index"`
	Memory int `gorm:"default:0"`
	RuntimeID uint 
	Code string `gorm:"type:text;not null"`
	Handler string `gorm:"type:varchar(100);not null"`
	Path string `gorm:"type:varchar(100);not null"`
	Image string `gorm:"type:varchar(50);not null"`
	AccountId string `gorm:"type:varchar(100);not null"`
}  
 
type Runtime struct { 
	gorm.Model 
	Image string
	Name string
	Extension string
	Activator string
	Version string
	Functions []Function `gorm:"foreignKey:RuntimeID"`
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