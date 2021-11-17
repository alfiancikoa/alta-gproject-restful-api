package models

type Province struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	Name string `gorm:"type:varchar(50);not null" json:"name" form:"name"`
}
