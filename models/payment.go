package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID        int     `gorm:"primarykey; AUTO_INCREMENT" json:"id" form:"id"`
	Type      string  `gorm:"type:varchar(100);not null" json:"type" form:"type"`
	Name      string  `gorm:"type:varchar(100);not null" json:"name" form:"name"`
	Orders    []Order `gorm:"foreignKey:Payment_ID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
