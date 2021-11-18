package database

import (
	"alte/e-commerce/config"
	"alte/e-commerce/models"
)

// Query Get All Products
func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// Query Get All my Products
func GetMyProducts(user_id int) ([]models.Product, error) {
	var products []models.Product
	if err := config.DB.Where("user_id=?", user_id).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductByID(id int) (*models.Product, error) {
	var product models.Product
	tx := config.DB.Find(&product, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected > 0 {
		return &product, nil
	}
	return nil, nil
}

func GetProductByCategory(name string) (models.Product, error) {
	product := models.Product{}
	if err := config.DB.Where("category = ?", name).Find(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func CreateProduct(product models.Product) (models.Product, error) {
	if err := config.DB.Save(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func UpdateProduct(newProduct *models.Product, Id int, user_id int) (*models.Product, error) {
	product := models.Product{}
	tx := config.DB.Where("user_id=?", user_id).Find(&product, Id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	product.Title = newProduct.Title
	product.Desc = newProduct.Desc
	product.Price = newProduct.Price
	product.Status = newProduct.Status

	if tx.RowsAffected > 0 {
		if err := config.DB.Save(&product).Error; err != nil {
			return nil, err
		} else {
			return &product, nil
		}
	}
	return nil, nil
}

func DeleteProduct(id, user_id int) (*models.Product, error) {
	product := models.Product{}
	tx := config.DB.Where("id = ? AND user_id=?", id, user_id).Delete(&product)
	if err := tx.Error; err != nil {
		return nil, err
	}
	if tx.RowsAffected > 0 {
		return &product, nil
	}
	return nil, nil
}
