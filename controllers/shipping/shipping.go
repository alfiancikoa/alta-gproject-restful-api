package shipping

import (
	"alte/e-commerce/lib/database"
	"alte/e-commerce/models"
	"alte/e-commerce/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

// #-------------------------------------------------
// # Controller Table Shipp Type
// #-------------------------------------------------

// Function Create New Shipping Type Controller
func CreateShipTypeController(c echo.Context) error {
	newShipType := PostShipType{}
	if err := c.Bind(&newShipType); err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse())
	}
	shipType := models.Ship_Type{
		Name: newShipType.Name,
	}
	respon, err := database.ShipTypeInsert(shipType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "Success create a new shipping_type",
		"data":    respon,
	})
}

// Function Show All ShipTypes
func GetShipTypeController(c echo.Context) error {
	respon, err := database.ShipTypeGet()
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get all ship type", "data": respon,
	})
}

// #-------------------------------------------------
// # Controller Table Shipping
// #-------------------------------------------------

// Function Create New Shipping Type Controller
func CreateShippingController(c echo.Context) error {
	newShipping := PostShipping{}
	if err := c.Bind(&newShipping); err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse())
	}
	shipping := models.Shipping{
		Name:        newShipping.Name,
		Cost:        newShipping.Cost,
		ShipType_ID: uint(newShipping.ShipType_ID),
	}
	respon, err := database.ShippingInsert(shipping)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "Success create a new shipping_type",
		"data":    respon,
	})
}

// Function Show All Shipping
func GetShippingController(c echo.Context) error {
	respon, err := database.ShippingTpingGet()
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get all ship type", "data": respon,
	})
}
