package mongoapi

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/clientopt"
	"time"
)

func Demo1() {
	var (
		client     *mongo.Client
		err        error
		database   *mongo.Database
		collection *mongo.Collection
	)

	// 1,建立连接
	if client, err = mongo.Connect(context.TODO(), "mongodb://localhost:27018", clientopt.ConnectTimeout(5*time.Second)); err != nil {
		fmt.Println(err)
		return
	}

	// 2,选择数据库test
	database = client.Database("test")

	// 3,选择表my_collection
	collection = database.Collection("my_collection")

	collection = collection

}
