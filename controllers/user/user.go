package user

import (
	"alte/e-commerce/lib/database"
	"alte/e-commerce/models"
	"alte/e-commerce/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Function Get All user Controller
func GetAllUsersController(c echo.Context) error {
	respon, err := database.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get all user", "data": respon,
	})
}

// Function Get All User By ID Controller
func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.InvalidFormatMethodInput())
	}
	user, err := database.GetUserId(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}
	if user == nil {
		return c.JSON(http.StatusNotFound, responses.DataNotExist())
	}
	respon := GetUserResponse{
		Name:        user.Name,
		Email:       user.Email,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		Gender:      user.Gender,
		Birth:       user.Birth,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "Success get user",
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
	user := models.User{
		Name:        newUser.Name,
		Email:       newUser.Email,
		Password:    newUser.Password,
		PhoneNumber: newUser.PhoneNumber,
		Gender:      newUser.Gender,
		Birth:       newUser.Birth,
	}
	respon, err := database.InsertUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "Success create a new user",
		"data":    respon,
	})
}

// Function Edit User By ID Controller
func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.InvalidFormatMethodInput())
	}
	userRequest := EditUserRequest{}
	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse())
	}
	if userRequest.Name == "" || userRequest.Email == "" || userRequest.Password == "" {
		return c.JSON(http.StatusBadRequest, responses.InvalidFormatMethodInput())
	}
	user := models.User{
		Name:        userRequest.Name,
		Email:       userRequest.Email,
		Password:    userRequest.Password,
		PhoneNumber: userRequest.PhoneNumber,
		Gender:      userRequest.Gender,
		Birth:       userRequest.Birth,
	}
	respon, err := database.EditUser(&user, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}
	if respon == nil {
		return c.JSON(http.StatusNotFound, responses.DataNotExist())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "Success edit user",
		"data":    respon,
	})
}

// Function Delete User By ID Controller
func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.InvalidFormatMethodInput())
	}
	respon, e := database.DeleteUser(id)
	if e != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}
	if respon == nil {
		return c.JSON(http.StatusNotFound, responses.DataNotExist())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "Success deleted user",
	})
}
