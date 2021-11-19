package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          int    `gorm:"primarykey; AUTO_INCREMENT" json:"id" form:"id"`
	Title       string `gorm:"type:varchar(255); not null" json:"title" form:"title"`
	Desc        string `gorm:"type:varchar(500); not null" json:"desc" form:"desc"`
	Price       int    `gorm:"type:int(10); not null" json:"price" form:"price"`
	Status      string `gorm:"type:enum('ready', 'sold out')" json:"status" form:"status"`
	Category_ID int
	User_ID     int
	Photos      []Photo_Product `gorm:"foreignKey:Product_ID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
