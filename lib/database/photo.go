package database

import (
	"alte/e-commerce/config"
	"alte/e-commerce/models"
)

func GetAllPhoto() ([]models.Photo_Product, error) {
	var photo []models.Photo_Product
	if err := config.DB.Find(&photo).Error; err != nil {
		return photo, err
	}
	return photo, nil
}

func GetPhotoById(id int) (*models.Photo_Product, error) {
	photo := models.Photo_Product{}
	tx := config.DB.Find(&photo, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected > 0 {
		return &photo, nil
	}
	return nil, nil
}

func InsertPhoto(photo models.Photo_Product) (*models.Photo_Product, error) {
	tx := config.DB.Save(&photo)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &photo, nil
}

func UpdatePhoto(newPhoto *models.Photo_Product, Id int) (*models.Photo_Product, error) {
	photo := models.Photo_Product{}
	tx := config.DB.Find(&photo, Id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	photo.Photo_Title = newPhoto.Photo_Title
	photo.Size_Photo = newPhoto.Size_Photo
	photo.Resolution_Photo = newPhoto.Resolution_Photo
	if tx.RowsAffected > 0 {
		if err := config.DB.Save(&photo).Error; err != nil {
			return nil, err
		} else {
			return &photo, nil
		}
	}
	return nil, nil
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
