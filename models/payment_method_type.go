package models

type PaymentMethodType struct {
	ID             uint   `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	PayMethod_Type string `gorm:"type:varchar(50);not null" json:"paymethod_type" form:"paymethod_type"`
}
