package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	uri := os.Args[1]
	if uri == "" {
		panic("uri can not be blank.")
	}
	fmt.Println("mongo uri: ", uri)
	// 创建一个新的客户端并连接到 MongoDB
	// mongo.Connect()

	// 设置一个 10 秒的连接超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opt := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		log.Fatal(err)
	}

	// 确保在程序退出时断开连接
	defer client.Disconnect(ctx)

	// 查询所有数据库信息
	databases, err := client.ListDatabaseNames(ctx, bson.D{{"empty", false}})
	if err != nil {
		log.Fatal(err)
	}

	// 打印所有数据库名称
	fmt.Println("Databases:")
	for _, dbName := range databases {
		fmt.Println(" -", dbName)
	}
	fmt.Println("mongodb connected.")
}
