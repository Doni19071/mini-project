package appModel

import (
	"gorm.io/gorm"
)

type SeatDbModel struct {
	db *gorm.DB
}

func NewSeatDbModel(db *gorm.DB) *SeatDbModel {
	db.AutoMigrate(&Seat{})
	return &SeatDbModel{
		db: db,
	}
}

func (pm *SeatDbModel) GetAllSeat() ([]Seat, error) {
	var allSeat []Seat
	err := pm.db.Find(&allSeat).Error
	return allSeat, err
}

func (pm *SeatDbModel) AddSeat(p Seat) (Seat, error) {
	err := pm.db.Save(&p).Error
	return p, err
}
