package models

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID        int        `gorm:"primarykey; AUTO_INCREMENT" json:"id" form:"id"`
	User_ID   int        `gorm:"primarykey" json:"user_id" form:"user_id"`
	User      User       `gorm:"foreignkey:User_ID;" json:"-"`
	CartItems []CartItem `gorm:"foreignKey:Cart_ID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
