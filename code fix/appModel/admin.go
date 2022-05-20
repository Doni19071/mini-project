package appModel

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Usernameame string `json:"name"`
	Password    string `json:"password"`
	Token       string `json:"token"`
}
