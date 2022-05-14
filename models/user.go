package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	//	Orders []Order `gorm:"foreignKey:UserRefer"`
}

type UserResponse struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
	Token    string `json:"token" form:"token"`
}
