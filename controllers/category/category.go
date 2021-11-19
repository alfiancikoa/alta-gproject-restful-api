package category

import (
	"alte/e-commerce/lib/database"
	"alte/e-commerce/models"
	"alte/e-commerce/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func InsertCategoryController(c echo.Context) error {
	newCategory := PostCategory{}
	if err := c.Bind(&newCategory); err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse())
	}

	if newCategory.Title == "" {
		return c.JSON(http.StatusBadRequest, responses.InvalidFormatMethodInput())
	}
	category := models.Category{
		Title: newCategory.Title,
	}

	respon, err := database.InsertCategory(category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "success",
		"data":   respon,
	})
}

//find all category
func GetAllCategorysController(c echo.Context) error {
	categorys, err := database.GetAllCategory()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success get all category",
		"data":   categorys,
	})
}

func UpdateCategoryController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		return c.JSON(http.StatusBadRequest, responses.InvalidFormatMethodInput())
	}
	categoryRequest := EditCategory{}
	if err := c.Bind(&categoryRequest); err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse())
	}
	if categoryRequest.Title == "" {
		return c.JSON(http.StatusBadRequest, responses.InvalidFormatMethodInput())
	}
	category := models.Category{
		Title: categoryRequest.Title,
	}
	respon, err := database.EditCategory(&category, id)
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

func DeleteCategoryController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		return c.JSON(http.StatusBadRequest, responses.InvalidFormatMethodInput())
	}

	respon, err := database.DeleteCategory(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}
	if respon == nil {
		return c.JSON(http.StatusNotFound, responses.DataNotExist())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "category succesfully deleted",
	})
}
