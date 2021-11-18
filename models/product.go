package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint   `gorm:"primarykey; AUTO_INCREMENT" json:"id" form:"id"`
	Title       string `gorm:"type:varchar(255); not null" json:"title" form:"title"`
	Desc        string `gorm:"type:varchar(500); not null" json:"desc" form:"desc"`
	Price       uint   `gorm:"type:int(10); not null" json:"price" form:"price"`
	Status      string `gorm:"type:enum('ready', 'sold out')" json:"status" form:"status"`
	Category_ID uint
	User_ID     uint
	Photos      []Photo_Product `gorm:"foreignKey:Product_ID;references:ID"`
	Create_at   time.Time
	Update_at   time.Time
	Delete_at   gorm.DeletedAt `gorm:"index"`
}
