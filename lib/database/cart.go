package database

import (
	"alte/e-commerce/config"
	"alte/e-commerce/models"
)

// Query Add Cart Item
func AddCartItem(item models.CartItem, cart_id int) (*models.CartItem, error) {
	tx := config.DB.Where("product_id=? AND cart_id=? and order_id=0", item.Product_ID, cart_id).Find(&models.CartItem{})
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected < 1 {
		if err := config.DB.Save(&item).Error; err != nil {
			return nil, err
		} else {
			return &item, nil
		}
	}
	return nil, nil
}

//Update Quantity
func UpdateQuantityCartItem(cartID int, item *models.CartItem) (*models.CartItem, error) {
	var cartItem models.CartItem
	if err := config.DB.Where("cart_id = ? and product_id = ? and order_id=0", cartID, item.Product_ID).First(&cartItem).Error; err != nil {
		return &cartItem, err
	}
	updatedQuantity := item.Qty + cartItem.Qty
	updatePrice := item.Price + cartItem.Price
	if err := config.DB.Model(&cartItem).Update("qty", updatedQuantity).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Model(&cartItem).Update("price", updatePrice).Error; err != nil {
		return nil, err
	}
	return &cartItem, nil
}

// Query Is Product Already Exist
func ProductAlreadyExist(id int) int {
	var product models.Product
	row := config.DB.Where("id = ?", id).Find(&product).RowsAffected
	return int(row)
}

// Query Get Cart_ID
func GetCartId(user_id int) (models.Cart, error) {
	var cart models.Cart
	if err := config.DB.Where("user_id=?", user_id).Find(&cart).Error; err != nil {
		return cart, err
	}
	return cart, nil
}

// Query Get Price Product
func GetPriceProduct(product_id int) (int, error) {
	product := models.Product{}
	tx := config.DB.Where("id=?", product_id).Find(&product)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return product.Price, nil
}

// Query Get myCart
func GetmyCart(cart_id int) ([]models.CartItem, error) {
	cart := []models.CartItem{}
	if err := config.DB.Where("cart_id=? and status='0'", cart_id).Find(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

// Query Update Cart
func EditCart(newItem *models.CartItem, Id int, cart_id int) (*models.CartItem, error) {
	item := models.CartItem{}
	tx := config.DB.Where("cart_id=?", cart_id).Find(&item, Id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	price, _ := GetPriceProduct(item.Product_ID)
	item.Qty = newItem.Qty
	item.Price = price * newItem.Qty
	if tx.RowsAffected > 0 {
		if err := config.DB.Save(&item).Error; err != nil {
			return nil, err
		} else {
			return &item, nil
		}
	}
	return nil, nil
}

// Query Delete Cart
func DeleteCart(id, cart_id int) (*models.CartItem, error) {
	tx := config.DB.Where("id=? AND cart_id=?", id, cart_id).Delete(&models.CartItem{})
	if err := tx.Error; err != nil {
		return nil, err
	}
	if tx.RowsAffected > 0 {
		return &models.CartItem{}, nil
	}
	return nil, nil
}

// Query Creat Cart in User Id
func CreateCart(cart models.Cart) error {
	if err := config.DB.Save(&cart).Error; err != nil {
		return err
	}
	return nil
}
