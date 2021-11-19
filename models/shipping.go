package models

import (
	"time"

	"gorm.io/gorm"
)

type Ship_Type struct {
	ID        int        `gorm:"primarykey"`
	Name      string     `gorm:"type:varchar(100);not null" json:"name" form:"name"`
	Shippings []Shipping `gorm:"foreignKey:ShipType_ID;references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
type Shipping struct {
	ID          int    `gorm:"primarykey"`
	Name        string `gorm:"type:varchar(100);not null" json:"name" form:"name"`
	Cost        int    `gorm:"type:int;not null" json:"cost" form:"cost"`
	ShipType_ID int    `json:"shiptype_id" form:"shiptype_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
