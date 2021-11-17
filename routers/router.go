package routers

import (
	"alte/e-commerce/controllers/user"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	// ------------------------------------------------------------------
	// USERS ROUTER
	// ------------------------------------------------------------------
	e.GET("/users", user.GetAllUsersController)
	e.GET("/users/:id", user.GetUserController)
	e.POST("/users", user.CreateUserController)
	e.PUT("/users/:id", user.UpdateUserController)
	e.DELETE("/users/:id", user.DeleteUserController)

	return e
}
