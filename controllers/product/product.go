package product

import (
	"alte/e-commerce/lib/database"
	"alte/e-commerce/middlewares"
	"alte/e-commerce/models"
	"alte/e-commerce/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateProductsController(c echo.Context) error {
	newProduct := PostProduct{}
	if err := c.Bind(&newProduct); err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse())
	}
	if newProduct.Title == "" || newProduct.Desc == "" || newProduct.Price <= 0 || newProduct.Stock < 0 || newProduct.Category_ID <= 0 {
		return c.JSON(http.StatusBadRequest, responses.InvalidFormatMethodInput())
	}
	product := models.Product{
		Title:       newProduct.Title,
		Desc:        newProduct.Desc,
		Price:       newProduct.Price,
		Stock:       newProduct.Stock,
		Category_ID: newProduct.Category_ID,
	}
	product.User_ID = middlewares.ExtractTokenUserId(c)
	_, err := database.CreateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success create new product",
	})
}

//find all product
func GetAllProductsController(c echo.Context) error {
	products, err := database.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}
	respon := make([]GetProductResponse, len(products))
	for i := 0; i < len(products); i++ {
		respon[i].ID = products[i].ID
		respon[i].Title = products[i].Title
		respon[i].Desc = products[i].Desc
		respon[i].Price = products[i].Price
		respon[i].Stock = products[i].Stock
		respon[i].Category_ID = products[i].Category_ID
		respon[i].User_ID = products[i].User_ID
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success get all product",
		"data":    respon,
	})
}

// Get All My Product
func GetMyProductController(c echo.Context) error {
	user_id := middlewares.ExtractTokenUserId(c)
	products, err := database.GetMyProducts(user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}
	respon := make([]GetProductResponse, len(products))
	for i := 0; i < len(products); i++ {
		respon[i].ID = products[i].ID
		respon[i].Title = products[i].Title
		respon[i].Desc = products[i].Desc
		respon[i].Price = products[i].Price
		respon[i].Stock = products[i].Stock
		respon[i].Category_ID = products[i].Category_ID
		respon[i].User_ID = products[i].User_ID
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success get all product",
		"data":    respon,
	})
}

func GetProductController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		return c.JSON(http.StatusBadRequest, responses.InvalidFormatMethodInput())
	}
	product, err := database.GetProductByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())

	}
	if product == nil {
		return c.JSON(http.StatusNotFound, responses.DataNotExist())
	}
	respons := GetProductResponse{
		ID:          product.ID,
		Title:       product.Title,
		Desc:        product.Desc,
		Price:       product.Price,
		Stock:       product.Stock,
		Category_ID: product.Category_ID,
		User_ID:     product.User_ID,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success get product",
		"data":    respons,
	})
}

func UpdateProductController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		return c.JSON(http.StatusBadRequest, responses.InvalidFormatMethodInput())
	}
	productRequest := EditProduct{}
	if err := c.Bind(&productRequest); err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse())
	}
	if productRequest.Title == "" || productRequest.Desc == "" || productRequest.Price <= 0 || productRequest.Stock < 0 {
		return c.JSON(http.StatusBadRequest, responses.InvalidFormatMethodInput())
	}
	product := models.Product{
		Title:       productRequest.Title,
		Desc:        productRequest.Desc,
		Price:       productRequest.Price,
		Stock:       productRequest.Stock,
		Category_ID: productRequest.Category_ID,
	}
	user_id := middlewares.ExtractTokenUserId(c)
	respon, err := database.UpdateProduct(&product, id, user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}
	if respon == nil {
		return c.JSON(http.StatusNotFound, responses.DataNotExist())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "product update successful",
	})
}

func DeleteProductController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, responses.InvalidFormatMethodInput())
	}
	user_id := middlewares.ExtractTokenUserId(c)
	respon, err := database.DeleteProduct(id, user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}
	if respon == nil {
		return c.JSON(http.StatusNotFound, responses.DataNotExist())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success", "message": "product deleted successfully",
	})
}

// Testing Get User
func GetMyProductControllerTesting() echo.HandlerFunc {
	return GetMyProductController
}

// Testing Update User
func UpdateProductControllerTesting() echo.HandlerFunc {
	return UpdateProductController
}

// Testing Delete User
func DeleteProductControllerTesting() echo.HandlerFunc {
	return DeleteProductController
}

// Testing Create User
func CreateProductsControllerTesting() echo.HandlerFunc {
	return CreateProductsController
}
