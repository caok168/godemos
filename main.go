package main

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"godemos/influxdbdemo"
	"godemos/mysql"
	"godemos/postgres"
	"godemos/postgres/models"
)

var (
	pgDbUrl = "host=localhost port=5431 user=test password=test dbname=test sslmode=disable"
	//mysqlDbUrl = "root:test@tcp(mysql:3307)/test"
	mysqlDbUrl = "root:test@tcp(mysql:3307)/test"
)

func main() {
	fmt.Println("test")

	//postgresDemo()
	//mysqlDemo()
	//mongoDemo()
	//mongoapi.Demo1()
	//mongoapi.Demo2()
	//mongoapi.Demo3()
	//mongoapi.Demo4()
	//mongoapi.Demo5()

	// casbin
	//casbindemo.Demo("postgres", pgDbUrl)
	//casbindemo.Demo2()
	//casbindemo.Demo3()

	//influxdbdemo.Demo()
	influxdbdemo.Demo2()
}

// postgres
func postgresDemo() {
	db, err := postgres.ConnectPostgres(pgDbUrl, false)
	if err != nil {
		fmt.Println("err :", err)
		return
	} else {
		fmt.Println("db:", db)
	}

	db = db.AutoMigrate(&models.Student{}).AutoMigrate(&models.Setting{})
}

// mysql
func mysqlDemo() {

	db, err := mysql.ConnectMySQL(mysqlDbUrl, false)
	if err != nil {
		fmt.Println("err: ", err)
		return
	} else {
		fmt.Println("db:", db)
	}

	db = db.AutoMigrate(&models.Student{}).AutoMigrate(&models.Setting{})
}

// mongo
func mongoDemo() {
	client, err := mongo.Connect(context.Background(), "mongodb://localhost:27018/test", nil)
	if err != nil {
		panic(err)
	}

	db := client.Database("test")
	fmt.Println(db)
}
