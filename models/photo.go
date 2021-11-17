package models

import "time"

type Photo_Product struct {
	ID_Photo         int    `gorm:"primarykey;AUTO_INCREMENT" json:"id" form:"id"`
	Photo_Title      string `gorm:"type:varchar(255);unique;not null" json:"title" form:"title"`
	Size_Photo       string `gorm:"type:varchar(15);not null" json:"size" form:"size"`
	Resolution_Photo string `gorm:"type:varchar(15);not null" json:"resolution" form:"resolution"`
	Link_Photo       string `gorm:"type:url" json:"link" form:"link"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
