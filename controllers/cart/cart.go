package cart

import (
	db "alte/e-commerce/lib/database"
	"alte/e-commerce/middlewares"
	"alte/e-commerce/models"
	"alte/e-commerce/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateCartController(c echo.Context) error {
	cartItem := models.CartItem{}
	user_id := middlewares.ExtractTokenUserId(c)
	if err := c.Bind(&cartItem); err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse())
	}
	isExist := db.ProductAlreadyExist(cartItem.Product_ID)
	if isExist == 0 {
		return c.JSON(http.StatusBadGateway, responses.DataNotExist())
	}
	cart, _ := db.GetCartId(user_id)
	productPrice, _ := db.GetPriceProduct(cartItem.Product_ID)
	totprice := productPrice * cartItem.Qty
	cartItem.Price = totprice
	cartItem.Cart_ID = cart.ID
	respon, err := db.AddCartItem(cartItem, cart.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}
	if respon == nil {
		_, err := db.UpdateQuantityCartItem(cart.ID, &cartItem)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.DataAlreadyExist())
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success create new cart",
	})
}

func GetCartController(c echo.Context) error {
	user_id := middlewares.ExtractTokenUserId(c)
	carts, err := db.GetmyCart(user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}
	respon := make([]GetCartRespon, len(carts))
	for i := 0; i < len(carts); i++ {
		respon[i].ID = carts[i].ID
		respon[i].Product_ID = carts[i].Product_ID
		respon[i].Qty = carts[i].Qty
		respon[i].Price = carts[i].Price
		respon[i].User_ID = user_id
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success get my cart",
		"data":    respon,
	})
}

func UpdateCartController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, responses.InvalidFormatMethodInput())
	}
	cartRequest := EditCartItemRequest{}
	if err := c.Bind(&cartRequest); err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse())
	}
	cartItem := models.CartItem{
		Qty: cartRequest.Qty,
	}
	user_id := middlewares.ExtractTokenUserId(c)
	respon, err := db.EditCart(&cartItem, id, user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}
	if respon == nil {
		return c.JSON(http.StatusNotFound, responses.DataNotExist())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "cart update successful",
	})
}

func DeleteCartController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, responses.InvalidFormatMethodInput())
	}
	user_id := middlewares.ExtractTokenUserId(c)
	respon, err := db.DeleteCart(id, user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}
	if respon == nil {
		return c.JSON(http.StatusNotFound, responses.DataNotExist())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success", "message": "cart deleted successfully",
	})
}
