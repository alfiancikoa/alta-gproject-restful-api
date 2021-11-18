package photo

import (
	"alte/e-commerce/lib/database"
	"alte/e-commerce/models"
	"alte/e-commerce/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func InsertPhotoController(c echo.Context) error {
	newPhoto := PostPhoto{}
	if err := c.Bind(&newPhoto); err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse())
	}

	if newPhoto.Title == "" || newPhoto.Size == "" || newPhoto.Resolution == "" || newPhoto.Link == "" || newPhoto.Product_ID <= 0 {
		return c.JSON(http.StatusBadRequest, responses.InvalidFormatMethodInput())
	}
	photo := models.Photo_Product{
		Title:      newPhoto.Title,
		Size:       newPhoto.Size,
		Resolution: newPhoto.Resolution,
		Link:       newPhoto.Link,
		Product_ID: newPhoto.Product_ID,
	}

	respon, err := database.InsertPhoto(photo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "success",
		"data":   respon,
	})
}

//find all photo
func GetAllPhotosController(c echo.Context) error {
	photos, err := database.GetAllPhoto()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success get all photo",
		"data":   photos,
	})
}

func DeletePhotoController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		return c.JSON(http.StatusBadRequest, responses.InvalidFormatMethodInput())
	}

	respon, err := database.DeletePhoto(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}
	if respon == nil {
		return c.JSON(http.StatusNotFound, responses.DataNotExist())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "photo succesfully deleted",
	})
}
