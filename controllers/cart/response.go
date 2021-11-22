package cart

type AddCartItem struct {
	Qty        int `json:"qty" form:"qty"`
	Product_ID int `json:"product_id" form:"product_id"`
	Price      int
}
type EditCartItemRequest struct {
	Qty        int    `json:"qty" form:"qty"`
	Price      int    `json:"price" form:"price"`
	Status     string `json:"status" form:"status"`
	Product_ID int    `json:"product_id" form:"product_id"`
	User_ID    int    `json:"user_id" form:"user_id"`
}

type GetCartRespon struct {
	ID         int
	Qty        int
	Price      int
	Product_ID int
	User_ID    int
}
