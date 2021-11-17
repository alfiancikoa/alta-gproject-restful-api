package models

type ShippingType struct {
	ID            uint   `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	Shipping_Type string `gorm:"type:varchar(50);not null" json:"shipping_type" form:"shipping_type"`
}
