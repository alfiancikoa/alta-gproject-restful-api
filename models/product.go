package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          int    `gorm:"primarykey; AUTO_INCREMENT" json:"id" form:"id"`
	Title       string `gorm:"type:varchar(255); not null" json:"title" form:"title"`
	Desc        string `gorm:"type:varchar(500); not null" json:"desc" form:"desc"`
	Price       int    `gorm:"type:int; not null" json:"price" form:"price"`
	Stock       int    `gorm:"type:int" json:"stock" form:"stock"`
	Category_ID int
	User_ID     int
	CartItems   []CartItem      `gorm:"foreignKey:Product_ID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Photos      []Photo_Product `gorm:"foreignKey:Product_ID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type ResponOrderProduct struct {
	ID    int
	Title string
	Desc  string
	Price int
	Stock int
}
