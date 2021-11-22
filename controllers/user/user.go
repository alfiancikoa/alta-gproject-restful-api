package user

import (
	"alte/e-commerce/lib/database"
	"alte/e-commerce/middlewares"
	"alte/e-commerce/models"
	"alte/e-commerce/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Login User Controller
func LoginUsersController(c echo.Context) error {
	var userlogin models.User
	if err := c.Bind(&userlogin); err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse())
	}
	user, err := database.GetUserByEmail(userlogin)
	if err != nil || user == nil {
		return c.JSON(http.StatusBadRequest, responses.InvalidEmailPassword())
	}
	respon, _ := database.GenerateToken(user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "login success", "data": respon.Token,
	})
}

// GET user by id User
func GetUserByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.InvalidFormatMethodInput())
	}
	loggedInUserId := middlewares.ExtractTokenUserId(c)
	if loggedInUserId != id {
		return c.JSON(http.StatusUnauthorized, responses.UnAuthorized())
	}
	users, err := database.GetUserId(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}
	respon := GetUserResponse{
		ID:          users.ID,
		Name:        users.Name,
		Email:       users.Email,
		PhoneNumber: users.PhoneNumber,
		Gender:      users.Gender,
		Birth:       users.Birth,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success get user detail",
		"data":    respon,
	})
}

// Function Create New User Controller
func CreateUserController(c echo.Context) error {
	newUser := PostUserRequest{}
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse())
	}
	if newUser.Name == "" || newUser.Email == "" || newUser.Password == "" {
		return c.JSON(http.StatusBadRequest, responses.InvalidFormatMethodInput())
	}
	xpass, _ := database.EncryptPassword(newUser.Password)
	user := models.User{
		Name:        newUser.Name,
		Email:       newUser.Email,
		Password:    xpass,
		PhoneNumber: newUser.PhoneNumber,
		Gender:      newUser.Gender,
		Birth:       newUser.Birth,
		Role:        "user",
	}
	userDetail, err := database.InsertUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}
	cart := models.Cart{User_ID: userDetail.ID}
	if e := database.CreateCart(cart); e != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success create a new user",
	})
}

// Function Edit User By ID Controller
func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.InvalidFormatMethodInput())
	}
	loggedInUserId := middlewares.ExtractTokenUserId(c)
	if loggedInUserId != id {
		return c.JSON(http.StatusUnauthorized, responses.UnAuthorized())
	}
	userRequest := EditUserRequest{}
	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse())
	}
	if userRequest.Name == "" || userRequest.Email == "" || userRequest.Password == "" {
		return c.JSON(http.StatusBadRequest, responses.InvalidFormatMethodInput())
	}
	xpass, _ := database.EncryptPassword(userRequest.Password)
	user := models.User{
		Name:        userRequest.Name,
		Email:       userRequest.Email,
		Password:    xpass,
		PhoneNumber: userRequest.PhoneNumber,
		Gender:      userRequest.Gender,
		Birth:       userRequest.Birth,
	}
	if _, err := database.EditUser(&user, id); err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success edit user",
	})
}

// Function Delete User By ID Controller
func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.InvalidFormatMethodInput())
	}
	loggedInUserId := middlewares.ExtractTokenUserId(c)
	if loggedInUserId != id {
		return c.JSON(http.StatusUnauthorized, responses.UnAuthorized())
	}
	_, e := database.DeleteUser(id)
	if e != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success deleted user",
	})
}

// Testing Get User
func GetUserByIdControllerTesting() echo.HandlerFunc {
	return GetUserByIdController
}

// Testing Edit User
func UpdateUserControllerTesting() echo.HandlerFunc {
	return UpdateUserController
}

// Testing Detele User
func DeleteUserControllerTesting() echo.HandlerFunc {
	return DeleteUserController
}
