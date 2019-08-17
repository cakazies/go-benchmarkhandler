package models

import "github.com/jinzhu/gorm"

type Users struct {
	Email      string `json:"email"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	Password   string `json:"password,omitemp"`
	LastName   string `json:"last_name,omitemp"`
	Address    string `json:"address,omitemp"`
	Zipcode    string `json:"zipcode,omitemp"`
	Gender     string `json:"gender,omitemp"`
	Phone      string `json:"phone,omitemp"`
	Avatar     string `json:"avatar,omitemp"`
	IdTypeUser string `json:"idtypeuser,omitemp"`
	gorm.Model
}

func GetProfile() {
	var value Users
	Conn.Where("id = ?", "1").First(&value)
}
