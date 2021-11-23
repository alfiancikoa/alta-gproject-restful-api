package order

import (
	db "alte/e-commerce/lib/database"
	"alte/e-commerce/middlewares"
	"alte/e-commerce/models"
	"alte/e-commerce/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateNewOrderController(c echo.Context) error {
	detailorder := models.PostOrderReq{}
	user_id := middlewares.ExtractTokenUserId(c)
	if err := c.Bind(&detailorder); err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequestResponse())
	}
	cart_items, err := db.GetCartItems(detailorder.CartItem_ID, user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}
	var totPrice, totQty int
	for i := 0; i < len(cart_items); i++ {
		totPrice += cart_items[i].Price
		totQty += cart_items[i].Qty
	}
	if totPrice < 1 || totQty < 1 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "failed", "message": "product belum tersedia di cart",
		})
	}
	detailorder.User_ID = user_id
	detailorder.Total_Price = totPrice
	detailorder.Total_Qty = totQty
	if _, err := db.NewOrder(detailorder); err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success create order",
	})
}

func GetOrderController(c echo.Context) error {
	user_id := middlewares.ExtractTokenUserId(c)
	orders, err := db.GetmyOrder(user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponse())
	}
	orderRespon := make([]OrderRespon, len(orders))
	for index, idOrder := range orders {
		idCartItems, _ := db.GetDataCartItems(idOrder.ID)
		cartRespon := make([]CartItemRespon, len(idCartItems))
		for i := 0; i < len(idCartItems); i++ {
			product, _ := db.GetProductByCartId(idCartItems[i].Product_ID)
			cartRespon[i].ID = idCartItems[i].ID
			cartRespon[i].Total_Qty = idCartItems[i].Qty
			cartRespon[i].Total_Price = idCartItems[i].Price
			cartRespon[i].Products = *product
		}
		orderRespon[index].Order_ID = idOrder.ID
		orderRespon[index].Total_Price = idOrder.Total_Price
		orderRespon[index].Total_Qty = idOrder.Total_Qty
		orderRespon[index].User_ID = idOrder.User_ID
		orderRespon[index].CartItem_ID = cartRespon
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success get my order",
		"data":    orderRespon,
	})
}
