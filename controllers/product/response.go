package product

type GetProductResponse struct {
	ID          int
	Title       string
	Desc        string
	Price       int
	Status      string
	Category_ID int
	User_ID     int
}

type PostProduct struct {
	Title       string `json:"title" form:"title"`
	Desc        string `json:"desc" form:"desc"`
	Price       int    `json:"price" form:"price"`
	Status      string `json:"status" form:"status"`
	Category_ID int    `json:"category_id" form:"category_id"`
}

type PostProductErr struct {
	Title int
}

type EditProduct struct {
	Title       string `json:"title" form:"title"`
	Desc        string `json:"desc" form:"desc"`
	Price       int    `json:"price" form:"price"`
	Status      string `json:"status" form:"status"`
	Category_ID int    `json:"category_id" form:"category_id"`
}

type EditProductErr struct {
	Title int
}
