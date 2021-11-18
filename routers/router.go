package routers

import (
	"alte/e-commerce/constants"
	"alte/e-commerce/controllers/category"
	"alte/e-commerce/controllers/product"
	"alte/e-commerce/controllers/user"

	"github.com/labstack/echo/v4"
	echoMid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	// ------------------------------------------------------------------
	// LOGIN & REGISTER USER
	// ------------------------------------------------------------------
	e.POST("/users", user.CreateUserController)
	e.POST("/login", user.LoginUsersController)
	// ------------------------------------------------------------------
	// JWT Authentication
	// ------------------------------------------------------------------
	eJWT := e.Group("")
	eJWT.Use(echoMid.JWT([]byte(constants.SECRET_JWT)))
	// ------------------------------------------------------------------
	// USERS ROUTER
	// ------------------------------------------------------------------
	eJWT.GET("/users/:id", user.GetUserByIdController)
	eJWT.PUT("/users/:id", user.UpdateUserController)
	eJWT.DELETE("/users/:id", user.DeleteUserController)
	// ------------------------------------------------------------------
	// CATEGORY & PRODUCT ROUTER
	// ------------------------------------------------------------------
	e.POST("/products/category", category.InsertCategoryController)
	e.GET("/products/category", category.GetAllCategorysController)
	eJWT.GET("/products", product.GetAllProductsController)
	eJWT.GET("/products/:id", product.GetProductController)
	eJWT.GET("/myproducts", product.GetMyProductController)
	eJWT.POST("/products", product.CreateProductsController)
	eJWT.DELETE("/products/:id", product.DeleteProductController)
	eJWT.PUT("/products/:id", product.UpdateProductController)
	return e
}
