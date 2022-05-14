package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	IdSeat string `json:"idSeat"`
	Name   string `json:"name"`
	NoHP   string `json:"noHP"`
}
