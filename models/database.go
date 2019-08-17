package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	Conn *gorm.DB
)

func init() {
	host := "127.0.0.1"
	port := "3306"
	user := "root"
	password := ""
	dbname := "crowde"

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbname)
	db, err := gorm.Open("mysql", conn)

	if err != nil {
		panic(err)
	}
	Conn = db
}
