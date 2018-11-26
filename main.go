package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	fmt.Println("test")
	dbURL := "host=localhost port=5431 user=test password=test dbname=test sslmode=disable"
	db, err := ConnectPostgres(dbURL, false)
	if err != nil {
		fmt.Println("err :", err)
		return
	} else {
		fmt.Println("db:", db)
	}

}

func ConnectPostgres(url string, extlog bool) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	db.LogMode(true)
	if extlog {
		//db.SetLogger()
	}

	db.SingularTable(false)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	return db, nil
}
