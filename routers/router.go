package routers

import (
	"alte/e-commerce/constants"
	"alte/e-commerce/controllers/shipping"
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
	// SHIPPING Routers
	// ------------------------------------------------------------------
	e.GET("/shipping/type", shipping.GetShipTypeController)
	e.POST("/shipping/type", shipping.CreateShipTypeController)
	e.GET("/shipping", shipping.GetShippingController)
	e.POST("/shipping", shipping.CreateShippingController)
	// ------------------------------------------------------------------
	// JWT Authentication
	// ------------------------------------------------------------------
	eJWT := e.Group("/jwt")
	eJWT.Use(echoMid.JWT([]byte(constants.SECRET_JWT)))
	// ------------------------------------------------------------------
	// USERS ROUTER
	// ------------------------------------------------------------------
	eJWT.GET("/users/:id", user.GetUserByIdController)
	eJWT.PUT("/users/:id", user.UpdateUserController)
	eJWT.DELETE("/users/:id", user.DeleteUserController)

	return e
}
