package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	IdSeat string `json:"idSeat"`
	// UserID uint
	Name string `json:"name"`
}
