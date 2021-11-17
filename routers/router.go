package routers

import (
	"alte/e-commerce/constants"
	"alte/e-commerce/controllers/user"

	"github.com/labstack/echo/v4"
	echoMid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	// ------------------------------------------------------------------
	// LOGIN Authentication REGISTER USER
	// ------------------------------------------------------------------
	e.POST("/users", user.CreateUserController)
	e.POST("/login", user.LoginUsersController)
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
