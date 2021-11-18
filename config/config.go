package config

import (
	"alte/e-commerce/models"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConfig() (config map[string]string) {
	conf, err := godotenv.Read()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return conf
}

var DB *gorm.DB

// Initial Database
func InitDB() {
	dbconfig := GetConfig()
	// Sesuaikan dengan database kalian
	connect := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		dbconfig["DB_USERNAME"],
		dbconfig["DB_PASSWORD"],
		dbconfig["DB_HOST"],
		dbconfig["DB_PORT"],
		dbconfig["DB_NAME"])

	var err error
	DB, err = gorm.Open(mysql.Open(connect), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitalMigration()
}

// Function Initial Migration
func InitalMigration() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Category{})
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.Ship_Type{})
	DB.AutoMigrate(&models.Shipping{})
}
