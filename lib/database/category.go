package database

import (
	"alte/e-commerce/config"
	"alte/e-commerce/models"
)

func GetAllCategory() ([]models.Category, error) {
	var category []models.Category
	if err := config.DB.Find(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

func GetCategoryById(id int) (*models.Category, error) {
	category := models.Category{}
	tx := config.DB.Find(&category, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected > 0 {
		return &category, nil
	}
	return nil, nil
}

func InsertCategory(category models.Category) (*models.Category, error) {
	tx := config.DB.Save(&category)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &category, nil
}

func DeleteCategory(Id int) (*models.Category, error) {
	category := models.Category{}
	tx := config.DB.Delete(&category, Id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected > 0 {
		return &category, nil
	}
	return nil, nil
}
