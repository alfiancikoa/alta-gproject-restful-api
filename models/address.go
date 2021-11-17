package models

type Address struct {
	ID             uint   `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	Detail_Address string `gorm:"type:varchar(200);not null" json:"detail_address" form:"detail_address"`
	Postal_Code    string `gorm:"type:varchar(5);not null" json:"postal_code" form:"postal_code"`
}
