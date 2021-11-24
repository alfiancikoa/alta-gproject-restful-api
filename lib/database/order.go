package database

import (
	"alte/e-commerce/config"
	"alte/e-commerce/models"
	"fmt"
)

func GetCartItem(cartItem_id int, cart_id int) (*models.CartItem, error) {
	cart := models.CartItem{}
	if err := config.DB.Where("id=? AND cart_id=? AND order_id=0", cartItem_id, cart_id).Find(&cart).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

func GetCartItems(id []int, cart_id int) ([]models.CartItem, error) {
	tampung := []models.CartItem{}
	for i := 0; i < len(id); i++ {
		cartItemId := id[i]
		respon, _ := GetCartItem(cartItemId, cart_id)
		tampung = append(tampung, *respon)
	}
	if len(tampung) > 0 {
		return tampung, nil
	}
	return nil, nil
}

func NewOrder(order models.PostOrderReq) (*models.Order, error) {
	neworder := models.Order{
		Total_Price: order.Total_Price,
		Total_Qty:   order.Total_Qty,
		User_ID:     order.User_ID,
		Payment_ID:  order.Payment_ID,
	}
	// Save Order Detail
	if err := config.DB.Save(&neworder).Error; err != nil {
		return nil, err
	}
	// Update Status Cart Item on the list order
	for i := 0; i < len(order.CartItem_ID); i++ {
		tx := config.DB.Model(&models.CartItem{}).Where("id=? AND order_id=0", order.CartItem_ID[i]).Updates(models.CartItem{Status: "1", Order_ID: neworder.ID})
		if err := tx.Error; err != nil {
			return nil, tx.Error
		}
	}
	// Genereate Transaction ID
	user, _ := GetUserId(neworder.User_ID)
	transaction_number := fmt.Sprintf("122%d%v", neworder.ID, user.PhoneNumber)
	if err := config.DB.Model(&models.Order{}).Where("id=?", neworder.ID).Update("transaction_number", transaction_number).Error; err != nil {
		return nil, err
	}
	// Create Addres reference by Order ID
	address := models.Address{
		Street:   order.Address.Street,
		City:     order.Address.City,
		State:    order.Address.State,
		Zip:      order.Address.Zip,
		Order_ID: neworder.ID,
	}
	if err := config.DB.Save(&address).Error; err != nil {
		return nil, err
	}
	// Return Transaction ID for Purchase the Order
	config.DB.Find(&neworder, neworder.ID)
	return &neworder, nil
}

func GetmyWaitingOrder(user_id int) ([]models.Order, error) {
	order := []models.Order{}
	if err := config.DB.Where("user_id=? and order_status = 'waiting'", user_id).Find(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func GetmyOrder(user_id int) ([]models.Order, error) {
	order := []models.Order{}
	if err := config.DB.Where("user_id=? and order_status = 'success' or order_status = 'waiting'", user_id).Find(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func GetDataCartItems(order_id int) ([]models.CartItem, error) {
	cartItem := []models.CartItem{}
	tx := config.DB.Where("order_id=?", order_id).Find(&cartItem)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return cartItem, nil
}

func GetProductByCartId(product_id int) (*models.ResponOrderProduct, error) {
	product := models.Product{}
	if err := config.DB.Where("id=?", product_id).Find(&product).Error; err != nil {
		return nil, err
	}
	responProduct := models.ResponOrderProduct{
		ID:    product.ID,
		Title: product.Title,
		Desc:  product.Desc,
		Price: product.Price,
		Stock: product.Stock,
	}
	return &responProduct, nil
}

func DeleteOrder(id, order_id int) (*models.Order, error) {
	if err := config.DB.Model(&models.Order{}).Where("id=?", id).Update("order_status", "cancell").Error; err != nil {
		return nil, err
	}
	tx := config.DB.Where("id=? AND user_id=?", id, order_id).Delete(&models.Order{})
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected > 0 {
		return &models.Order{}, tx.Error
	}
	return nil, nil
}

func GetMyCancelledOrders(user_id int) ([]models.Order, error) {
	order := []models.Order{}
	if err := config.DB.Unscoped().Where("user_id=? AND order_status='cancell'", user_id).Find(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func ConfirmSuccessOreder(transaction_number string) (*models.Order, error) {
	tx := config.DB.Model(&models.Order{}).Where("transaction_number=?", transaction_number).Update("order_status", "success")
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected > 0 {
		return &models.Order{}, nil
	}
	return nil, nil
}
