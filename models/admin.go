package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type AdminResponse struct {
	gorm.Model
	Username string `json:"username" form:"username"`
	Token    string `json:"token" form:"token"`
}
