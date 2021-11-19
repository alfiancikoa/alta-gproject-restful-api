package database

import (
	"alte/e-commerce/config"
	"alte/e-commerce/models"
)

func InsertPhoto(photo models.Photo_Product) (*models.Photo_Product, error) {
	tx := config.DB.Save(&photo)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &photo, nil
}

func GetAllPhoto() ([]models.Photo_Product, error) {
	var photo []models.Photo_Product
	if err := config.DB.Find(&photo).Error; err != nil {
		return photo, err
	}
	return photo, nil
}

func DeletePhoto(Id int) (*models.Photo_Product, error) {
	photo := models.Photo_Product{}
	tx := config.DB.Delete(&photo, Id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected > 0 {
		return &photo, nil
	}
	return nil, nil
}
