package models

import "gorm.io/gorm"

type Seat struct {
	gorm.Model
	Location string `json:"location"`
}
