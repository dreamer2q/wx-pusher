package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"wxServ/config"
)

var db *gorm.DB

const (
	dbDial = "%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
)

func Instance() *gorm.DB {
	return db
}

func Init() {
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf(dbDial,
		config.DB.User,
		config.DB.Pass,
		config.DB.Host,
		config.DB.DBName,
	))
	if err != nil {
		panic(err)
	}
}
