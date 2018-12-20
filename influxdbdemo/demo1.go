package influxdbdemo

import (
	"fmt"
	"github.com/influxdata/influxdb/client/v2"
	"log"
	"time"
)

const (
	MyDB     = "test"
	username = "root"
	password = "root"
)

func Demo() {
	conn, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://127.0.0.1:8086",
		Username: username,
		Password: password,
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(conn)

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})

	if err != nil {
		log.Fatal(err)
	}

	tags := map[string]string{"name": "xc"}
	fields := map[string]interface{}{
		"id":   1,
		"sex":  1,
		"pass": 0707,
	}

	pt, err := client.NewPoint("myuser", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	if err := conn.Write(bp); err != nil {
		log.Fatal(err)
	}

	fmt.Println("over")
}

// influxdb command
//show databases
//create database test
//show databases
//
//create user root with password 'root'
//grant all privileges on test to root
//
//show measurements