package appModel

import (
	"gorm.io/gorm"
)

type OrderDbModel struct {
	db *gorm.DB
}

func NewOrderDbModel(db *gorm.DB) *OrderDbModel {
	db.AutoMigrate(&Order{})
	return &OrderDbModel{
		db: db,
	}
}

func (pm *OrderDbModel) GetAllOrder() ([]Order, error) {
	var allOrder []Order
	err := pm.db.Find(&allOrder).Error
	return allOrder, err
}

func (pm *OrderDbModel) AddOrder(p Order) (Order, error) {
	err := pm.db.Save(&p).Error
	return p, err
}
