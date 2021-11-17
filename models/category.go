package models

import "time"

type Category struct {
	ID_Category    int    `gorm:"primarykey;AUTO_INCREMENT" json:"id" form:"id"`
	Category_Title string `gorm:"type:varchar(255);unique;not null" json:"title" form:"title"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
