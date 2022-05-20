package appModel

import (
	"gorm.io/gorm"
)

type AdminDBModel struct {
	db *gorm.DB
}

func NewAdminDBModel(db *gorm.DB) *AdminDBModel {
	db.AutoMigrate(&Admin{})
	return &AdminDBModel{
		db: db,
	}
}

func (pm *AdminDBModel) GetUsernameAndPassword(username string, password string) (Admin, error) {
	p := Admin{}
	err := pm.db.Where("username = ? AND password = ?", username, password).First(&p).Error
	return p, err
}

func (pm *AdminDBModel) AddAdmin(p Admin) (Admin, error) {
	err := pm.db.Save(&p).Error
	return p, err
}

func (pm *AdminDBModel) EditAdmin(id int, newP Admin) (Admin, error) {
	p := Admin{}
	// "select * from people where id=?", id
	err := pm.db.First(&p, id).Error
	if err != nil {
		return p, err
	}
	p.Usernameame = newP.Usernameame
	p.Password = newP.Password
	p.Token = newP.Token
	// "update person set ... where id=?", id
	err = pm.db.Save(&p).Error
	return p, err
}
