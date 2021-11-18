package database

import (
	"alte/e-commerce/config"
	"alte/e-commerce/models"
)

// #-------------------------------------------------
// # Table Shipp Type
// #-------------------------------------------------
// Query Create Shipp Type
func ShipTypeInsert(shipType models.Ship_Type) (*models.Ship_Type, error) {
	tx := config.DB.Save(&shipType)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &shipType, nil
}

// Query Select * From Shipp Type
func ShipTypeGet() (interface{}, error) {
	var shipType []models.Ship_Type
	if err := config.DB.Find(&shipType).Error; err != nil {
		return nil, err
	}
	return shipType, nil
}

// #-------------------------------------------------
// # Table Shipping
// #-------------------------------------------------
// Query Create Shipping
func ShippingInsert(shiping models.Shipping) (*models.Shipping, error) {
	tx := config.DB.Save(&shiping)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &shiping, nil
}

// Query Select * From Shipping
func ShippingGet() (interface{}, error) {
	var shipping []models.Shipping
	if err := config.DB.Find(&shipping).Error; err != nil {
		return nil, err
	}
	return shipping, nil
}

// Query Delete Sipping
func ShippingDelete(shippingId int) (*models.Shipping, error) {
	var shipping models.Shipping
	tx := config.DB.Delete(&shipping, shippingId)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected > 0 {
		return &shipping, nil
	}
	return nil, nil

}
