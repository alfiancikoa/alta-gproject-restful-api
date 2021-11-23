package order

import "alte/e-commerce/models"

type PostNewOrder struct {
	Total_Price int    `json:"total_price" form:"total_price"`
	Total_Qty   int    `json:"total_qty" form:"total_qty"`
	CartItem_ID []int  `json:"cartitem_id" form:"cartitem_id"`
	Payment_ID  int    `json:"payment_id" form:"payment_id"`
	User_ID     int    `json:"user_id" form:"user_id"`
	Address     string `json:"address" form:"address"`
}

type OrderRespon struct {
	Order_ID    int
	Total_Price int
	Total_Qty   int
	User_ID     int
	CartItem_ID []CartItemRespon
}
type CartItemRespon struct {
	ID          int
	Total_Qty   int
	Total_Price int
	Products    models.ResponOrderProduct
}
