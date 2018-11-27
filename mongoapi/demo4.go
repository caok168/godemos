package mongoapi

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/clientopt"
	"github.com/mongodb/mongo-go-driver/mongo/findopt"
	"time"
)

// jobName过滤条件
type FindByJobName struct {
	JobName string `bson:"jobName"`
}

func Demo4() {
	// mongodb 读取回来的是bson，需要反序列为LogRecord对象
	var (
		client     *mongo.Client
		err        error
		database   *mongo.Database
		collection *mongo.Collection
		cond       *FindByJobName
		cursor     mongo.Cursor
		record     *LogRecord
	)

	// 1, 建立连接
	if client, err = mongo.Connect(context.TODO(), "mongodb://localhost:27018", clientopt.ConnectTimeout(5*time.Second)); err != nil {
		fmt.Println(err)
		return
	}

	// 2, 选择数据库my_db
	database = client.Database("cron")

	// 3, 选择表my_collection
	collection = database.Collection("log")

	// 4, 按照jobName字段过滤，想找出jobName=job10,找出5条
	cond = &FindByJobName{JobName: "job10"}

	// 5, 查询（过滤 + 翻页参数）
	if cursor, err = collection.Find(context.TODO(), cond, findopt.Skip(0), findopt.Limit(2)); err != nil {
		fmt.Println(err)
		return
	}

	// 延迟释放游标
	defer cursor.Close(context.TODO())

	// 6, 遍历结果集
	for cursor.Next(context.TODO()) {
		// 定义一个日志对象
		record = &LogRecord{}

		// 反序列化bson到对象
		if err = cursor.Decode(record); err != nil {
			fmt.Println(err)
			return
		}

		// 把日志行打印出来
		fmt.Println(*record)
	}
}
