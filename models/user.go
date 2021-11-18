package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          int       `gorm:"primarykey"`
	Name        string    `gorm:"type:varchar(255)" json:"name" form:"name"`
	Email       string    `gorm:"type:varchar(100);unique;not null" json:"email" form:"email"`
	Password    string    `gorm:"type:varchar(255);not null" json:"password" form:"password"`
	PhoneNumber string    `gorm:"type:varchar(20);unique;not null" json:"phonenumber" form:"phonenumber"`
	Gender      string    `gorm:"type:enum('male','female');" json:"gender" form:"gender"`
	Birth       string    `gorm:"type:date" json:"birth" form:"birth"`
	Token       string    `gorm:"type:longtext;" json:"token" form:"token"`
	Role        string    `gorm:"type:enum('admin','user');"`
	Products    []Product `gorm:"foreignKey:User_ID;references:ID;constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
