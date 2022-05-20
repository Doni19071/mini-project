package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type AdminResponse struct {
	gorm.Model
	Username string `json:"username" form:"username"`
	Name     string `json:"name"`
	Token    string `json:"token" form:"token"`
}
