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
	if newProduct.Title == "" || newProduct.Desc == "" || newProduct.Price <= 0 || newProduct.Status == "" || newProduct.Category_ID <= 0 {
		return c.JSON(http.StatusBadRequest, responses.InvalidFormatMethodInput())
	}
	product := models.Product{
		Title:       newProduct.Title,
		Desc:        newProduct.Desc,
		Price:       newProduct.Price,
		Status:      newProduct.Status,
		Category_ID: newProduct.Category_ID,
	}
	product.User_ID = middlewares.ExtractTokenUserId(c)
	respon, err := database.CreateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "success",
		"message": "success create new product",
		"data":    respon,
	})
}

//find all product
func GetAllProductsController(c echo.Context) error {
	respon, err := database.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success get all product",
		"data":    respon,
	})
}

// Get All My Product
func GetMyProductController(c echo.Context) error {
	user_id := uint(middlewares.ExtractTokenUserId(c))
	respon, err := database.GetMyProducts(user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
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
		Title:  product.Title,
		Desc:   product.Desc,
		Price:  product.Price,
		Status: product.Status,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Success get user",
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
	if productRequest.Title == "" || productRequest.Desc == "" || productRequest.Price <= 0 || productRequest.Status == "" {
		return c.JSON(http.StatusBadRequest, responses.InvalidFormatMethodInput())
	}
	product := models.Product{
		Title:  productRequest.Title,
		Desc:   productRequest.Desc,
		Price:  productRequest.Price,
		Status: productRequest.Status,
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
		"status": "success",
		"data":   respon,
	})
}

func DeleteProductController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
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
		"status": "success", "message": "product succesfully deleted",
	})
}
