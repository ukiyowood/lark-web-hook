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

	// 查看所有的键
	fmt.Println("All keys in Redis:")
	var cursor uint64
	for {
		// 使用 SCAN 命令查找所有的键
		keys, nextCursor, err := rdb.Scan(ctx, cursor, "*", 0).Result()
		if err != nil {
			log.Fatalf("Could not scan keys: %v", err)
		}

		// 打印当前批次的键
		for _, key := range keys {
			fmt.Println(key)
		}

		// 如果 nextCursor 为 0，表示没有更多的键
		if nextCursor == 0 {
			break
		}
		cursor = nextCursor
	}
}
