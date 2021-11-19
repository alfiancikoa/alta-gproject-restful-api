package database

import (
	"alte/e-commerce/config"
	"alte/e-commerce/middlewares"
	"alte/e-commerce/models"

	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(userlogin *models.User) (*models.User, error) {
	user := models.User{}
	tx := config.DB.Where("id=?", userlogin.ID).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var err error
	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	config.DB.Save(user)
	return &user, nil
}

// Query Get User by Email
func GetUserByEmail(loginuser models.User) (*models.User, error) {
	user := models.User{}
	tx := config.DB.Where("email=?", loginuser.Email).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	checkpass := DecryptPassword(loginuser.Password, user.Password)
	if !checkpass {
		return nil, nil
	}
	return &user, nil
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

// Query Create New User Jwts
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

// Encrypt the password user
func EncryptPassword(password string) (string, error) {
	Encrypt, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return "", err
	}
	return string(Encrypt), nil
}

// Decrypt the password user
func DecryptPassword(password, encrypt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encrypt), []byte(password))
	return err == nil
}
