package mysql

import (
	"github.com/jinzhu/gorm"
	// import mysql driver for gorm
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConnectMySQL(url string, extlog bool) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", url+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	return db, nil
}
