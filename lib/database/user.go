package database

import (
	"alte/e-commerce/config"
	"alte/e-commerce/middlewares"
	"alte/e-commerce/models"
)

// Login User JWT
func Login(user *models.User) (*models.User, error) {
	tx := config.DB.Where("email=? AND password=?", user.Email, user.Password).First(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var err error
	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	if e := config.DB.Save(user).Error; e != nil {
		return nil, e
	}
	return user, nil
}

// Query Get user by Id
func GetUserId(userId int) (*models.User, error) {
	var user models.User
	tx := config.DB.Find(&user, userId)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected > 0 {
		return &user, nil
	}
	return nil, nil
}

// Query Create New User
func InsertUser(user models.User) (*models.User, error) {
	tx := config.DB.Save(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

// Query Edit User by Id
func EditUser(newUser *models.User, userId int) (*models.User, error) {
	user := models.User{}
	tx := config.DB.Find(&user, userId)
	if tx.Error != nil {
		return nil, tx.Error
	}
	user.Name = newUser.Name
	user.Email = newUser.Email
	user.Password = newUser.Password
	user.PhoneNumber = newUser.PhoneNumber
	user.Gender = newUser.Gender
	user.Birth = newUser.Birth
	if tx.RowsAffected > 0 {
		if err := config.DB.Save(&user).Error; err != nil {
			return nil, err
		} else {
			return &user, nil
		}
	}
	return nil, nil
}

// Query Delete User by Id
func DeleteUser(userId int) (*models.User, error) {
	var user models.User
	tx := config.DB.Delete(&user, userId)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected > 0 {
		return &user, nil
	}
	return nil, nil
}
