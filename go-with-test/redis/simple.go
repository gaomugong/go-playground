package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func example() {
	var ctx = context.Background()

	// get redis client from url
	redisUrl := "redis://localhost:6379?protocol=3"
	opts, err := redis.ParseURL(redisUrl)
	if err != nil {
		fmt.Printf("parse redis opts error: %s", err)
		panic(err)
	}
	redisClient := redis.NewClient(opts)

	// set and get key
	redisClient.Set(ctx, "hello", "world", time.Second*30)
	hello, err := redisClient.Get(ctx, "hello").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("redis ping: %v, hello=%s\n", redisClient.Ping(ctx), hello)

	// not exist key
	val, err := redisClient.Get(ctx, "not_exist_key").Result()
	if err == redis.Nil {
		fmt.Println("key not found")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Printf("val: %s", val)
	}

	// get redis client directly
	rdb := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
			// Default is 3.
			Protocol: 3,
		})

	// if err := rdb.Set(ctx, "key", "value", 0).Err(); err != nil {
	// 	panic(err)
	// }
	// fmt.Println(rdb.Del(ctx, "key"))
	if val, err := rdb.SetNX(ctx, "key", "value1", 10*time.Second).Result(); err != nil {
		fmt.Println(val, err)
	}
	fmt.Println(val, rdb.Get(ctx, "key").Err() == redis.Nil, rdb.Get(ctx, "key").Val())

	vals, err := rdb.Sort(ctx, "list", &redis.Sort{Offset: 0, Count: 2, Order: "ASC"}).Result()
	fmt.Println(vals, err)
}

func main() {
	example()
	lockTest()
}
