package routers

import (
	"alte/e-commerce/constants"
	"alte/e-commerce/controllers/cart"
	"alte/e-commerce/controllers/category"
	"alte/e-commerce/controllers/order"
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
	e.GET("/products", product.GetAllProductsController)
	eJWT.GET("/products/:id", product.GetProductController)
	eJWT.GET("/products/my", product.GetMyProductController)
	eJWT.POST("/products", product.CreateProductsController)
	eJWT.DELETE("/products/:id", product.DeleteProductController)
	eJWT.PUT("/products/:id", product.UpdateProductController)

	// ------------------------------------------------------------------
	// CART ROUTER
	// ------------------------------------------------------------------
	eJWT.POST("/carts", cart.CreateCartController)
	eJWT.GET("/carts/my", cart.GetCartController)
	eJWT.PUT("/carts/:id", cart.UpdateCartController)
	eJWT.DELETE("/carts/:id", cart.DeleteCartController)
	// ------------------------------------------------------------------
	// ORDER ROUTER
	// ------------------------------------------------------------------
	eJWT.POST("/orders", order.CreateNewOrderController)
	eJWT.GET("/orders", order.GetOrderController)
	eJWT.GET("/orders/history", order.GetOrderHistoryController)
	eJWT.DELETE("/orders/:id", order.CancellOrderController)
	e.POST("/orders/confirm", order.ConfirmOrderController)
	return e
}
