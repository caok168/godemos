package mongoapi

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/clientopt"
	"time"
)

func Demo3() {
	var (
		client     *mongo.Client
		err        error
		database   *mongo.Database
		collection *mongo.Collection
		record     *LogRecord
		logArr     []interface{}
		result     *mongo.InsertManyResult
		insertId   interface{}
		docId      objectid.ObjectID
	)

	// 1,建立连接
	if client, err = mongo.Connect(context.TODO(), "mongodb://localhost:27018", clientopt.ConnectTimeout(5*time.Second)); err != nil {
		fmt.Println(err)
		return
	}

	// 2,选择数据库my_db
	database = client.Database("cron")

	// 3,选择表my_collection
	collection = database.Collection("log")

	// 4,插入记录(bson)
	record = &LogRecord{
		JobName:   "job10",
		Command:   "echo hello",
		Err:       "",
		Content:   "hello",
		TimePoint: TimePoint{StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + 10},
	}

	// 5, 批量插入多条document
	logArr = []interface{}{record, record, record}

	// 发起插入
	if result, err = collection.InsertMany(context.TODO(), logArr); err != nil {
		fmt.Println(err)
		return
	}

	for _, insertId = range result.InsertedIDs {
		docId = insertId.(objectid.ObjectID)
		fmt.Println("自增ID：", docId.Hex())
	}

}
