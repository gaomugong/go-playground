package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	// 创建一个redis客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// 创建一个锁
	lock := NewLock(client, "my_lock")

	// 获取锁
	if lock.Acquire() {
		// 成功获取到锁，执行业务逻辑
		fmt.Println("Got lock, doing some work...")
		time.Sleep(time.Second * 5)
		fmt.Println("Work done, releasing lock...")

		// 释放锁
		lock.Release()
	} else {
		// 获取锁失败
		fmt.Println("Failed to get lock")
	}
}

// 分布式锁结构体
type Lock struct {
	client *redis.Client
	key    string
	value  string
}

// 创建一个新的分布式锁
func NewLock(client *redis.Client, key string) *Lock {
	return &Lock{
		client: client,
		key:    key,
		value:  "",
	}
}

// 获取锁
func (l *Lock) Acquire() bool {
	// 生成一个随机的value作为锁的值
	l.value = fmt.Sprintf("%d", time.Now().UnixNano())

	// 使用redis的setnx命令尝试获取锁
	success, err := l.client.SetNX(context.Background(), l.key, l.value, time.Second*10).Result()
	if err != nil {
		fmt.Println("Error acquiring lock:", err)
		return false
	}

	return success
}

// 释放锁
func (l *Lock) Release() bool {
	// 使用redis的lua脚本删除锁，保证删除的原子性，比如：
	// c1获取到锁，然后准备删除，且删除前判断锁值匹配，但此时恰好过期，并同时被c2获取到锁，这种情况下，c2的锁会被c1误删
	// 但是使用lua脚本，就能保证获取锁+判断+删除具备原子性
	script := `
        if redis.call("get", KEYS[1]) == ARGV[1] then
            return redis.call("del", KEYS[1])
        else
            return 0
        end
    `

	// 这种写法不可靠，故采用lua脚本实现
	// if ctx := context.Background(); l.client.Get(ctx, l.key).Val() == l.value {
	// 	l.client.Del(ctx, l.key)
	// }

	result, err := l.client.Eval(context.Background(), script, []string{l.key}, l.value).Result()
	if err != nil {
		fmt.Println("Error releasing lock:", err)
		return false
	}

	return result == int64(1)
}
