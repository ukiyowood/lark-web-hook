package main

import (
	"context"
	"fmt"
	"log"
	"os"

	redis "github.com/go-redis/redis/v8"
)

func main() {
	// 创建一个上下文
	ctx := context.Background()

	url := os.Args[1]
	pass := ""
	if len(os.Args) > 2 {
		pass = os.Args[2]
	}

	// 创建 Redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     url,  // Redis 服务器地址
		Password: pass, // Redis 密码，默认为空
		DB:       0,    // 默认数据库
	})

	// 测试连接
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println("Connected to Redis")
	// 关闭 Redis 客户端
	defer rdb.Close()
}
