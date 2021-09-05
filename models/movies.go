package models

import (
	"gorm.io/gorm"
)

type Movies struct {
	gorm.Model
	ID          int    `gorm:"column:id;type:varchar(255);not null; PRIMARY_KEY" json:"id"`
	Title       string `gorm:"column:title;type:varchar(255);not null" json:"title" validate:"required"`
	Slug        string `gorm:"column:slug;type:varchar(255);not null" json:"Slug" validate:"required"`
	Description string `gorm:"column:description;type:text" json:"Description" validate:"required"`
	Duration    uint   `gorm:"column:duration;type:int(5);not null" json:"Duration" validate:"required"`
	Image       string `gorm:"column:image;type:varchar(255);not null" json:"Image" validate:"required"`
	
}
