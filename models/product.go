package models

import "time"

type Product struct {
	ID_Product     uint          `gorm: "primarykey; AUTO_INCREMENT" json:"id" form:"id"`
	CategoryID     int           `json:"category_id" form:"category_id"`
	Category       Category      `gorm:"foreignkey:CategoryID" json:"-"`
	PhotoID        int           `json:"photo_id" form:"photo_id"`
	Photo          Photo_Product `gorm:"foreignkey:PhotoID" json:"-"`
	Product_Title  string        `gorm: "type:vatchar(255); not null" json:"title" form:"title"`
	Product_Desc   string        `gorm: "type:varchar(500); not null" json:"desc" form:"desc"`
	Product_Price  uint          `gorm: "type:int(10); not null" json:"price" form:"price"`
	Product_Status string        `gorm: "type:enum('sold', 'sold out')" json:"status" form:"status"`
	Create_at      time.Time
	Update_at      time.Time
}
