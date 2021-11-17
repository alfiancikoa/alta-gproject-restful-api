package database

import (
	"alte/e-commerce/config"
	"alte/e-commerce/models"
)

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductByID(id int) (models.Product, error) {
	var product models.Product
	if err := config.DB.Where("ID_Product = ?", id).Find(&product).Error; err != nil {
		return product, err
	}
	return product, nil
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

func UpdateProduct(id int, newProduct models.Product) (models.Product, error) {
	product := models.Product{}
	if err := config.DB.First(&product, id).Error; err != nil {
		return product, err
	}

	product.Product_Title = newProduct.Product_Title
	product.CategoryID = newProduct.CategoryID
	product.PhotoID = newProduct.PhotoID
	product.Product_Desc = newProduct.Product_Desc
	product.Product_Price = newProduct.Product_Price
	product.Product_Status = newProduct.Product_Status

	if err := config.DB.Save(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func DeleteProduct(id int) error {
	product := models.Product{}

	if err := config.DB.Where("id = ?", id).Delete(&product).Error; err != nil {
		return err
	}
	return nil
}
