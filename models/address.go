package models

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	ID        int    `gorm:"primarykey; AUTO_INCREMENT" json:"id" form:"id"`
	Street    string `gorm:"type:varchar(255);not null" json:"street" form:"street"`
	City      string `gorm:"type:varchar(255);not null" json:"city" form:"city"`
	State     string `gorm:"type:varchar(255);not null" json:"state" form:"state"`
	Zip       int    `gorm:"type:int;not null" json:"zip" form:"zip"`
	Order_ID  int
	Order     Order `gorm:"foreignKey:Order_ID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
