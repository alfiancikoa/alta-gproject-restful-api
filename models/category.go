package models

import "time"

type Category struct {
	ID        uint      `gorm:"primarykey;AUTO_INCREMENT" json:"id" form:"id"`
	Title     string    `gorm:"type:varchar(255);unique;not null" json:"title" form:"title"`
	Products  []Product `gorm:"foreignKey:Category_ID;references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
