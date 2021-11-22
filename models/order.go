package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID           int        `gorm:"primarykey; AUTO_INCREMENT" json:"id" form:"id"`
	Address      string     `gorm:"type:varchar(255);not null" json:"address" form:"address"`
	Total_Price  int        `gorm:"type:int;not null" json:"total_price" form:"total_price"`
	Total_Qty    int        `gorm:"type:int;not null" json:"total_qty" form:"total_qty"`
	Order_Status int        `gorm:"type:varchar;default:NULL" json:"order_status" form:"order_status"`
	User_ID      int        `json:"user_id" form:"user_id"`
	Payment_ID   int        `json:"payment_id" form:"payment_id"`
	CartItems    []CartItem `gorm:"foreignKey:Order_ID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type PostOrderReq struct {
	Total_Price int    `json:"total_price" form:"total_price"`
	Total_Qty   int    `json:"total_qty" form:"total_qty"`
	CartItem_ID []int  `json:"cartitem_id" form:"cartitem_id"`
	Payment_ID  int    `json:"payment_id" form:"payment_id"`
	User_ID     int    `json:"user_id" form:"user_id"`
	Address     string `json:"address" form:"address"`
}

type GetOrderReq struct {
	T_Price    int    `json:"total_price" form:"total_price"`
	T_Qty      int    `json:"total_qty" form:"total_qty"`
	Product_ID int    `json:"product_id" form:"product_id"`
	Title      int    `json:"title" form:"title"`
	Price      int    `json:"price" form:"price"`
	Stock      int    `json:"stock" form:"stock"`
	Desc       string `json:"desc" form:"desc"`
}
