package models

import (
	"time"

	"gorm.io/gorm"
)

type Photo_Product struct {
	ID         int    `gorm:"primarykey;AUTO_INCREMENT" json:"id" form:"id"`
	Title      string `gorm:"type:varchar(255);unique;not null" json:"title" form:"title"`
	Size       string `gorm:"type:varchar(15);not null" json:"size" form:"size"`
	Resolution string `gorm:"type:varchar(15);not null" json:"resolution" form:"resolution"`
	Link       string `gorm:"type:varchar(255)" json:"link" form:"link"`
	Product_ID int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Delete_at  gorm.DeletedAt `gorm:"index"`
}
