package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	lock    = "lock"
	timeout = time.Second * 30
	rdb     *redis.Client
	count   = 0
)

// Unlock 释放分布式锁
func Unlock(ctx context.Context) {
	rdb := getRdb()
	fmt.Printf("release lock: %v\n", rdb.Get(ctx, lock).Val())
	fmt.Println(rdb.Del(ctx, lock))

}

// Lock 获取分布式锁
func Lock(ctx context.Context) bool {
	rdb := getRdb()
	lockValue := rand.Int()
	fmt.Printf("acquire lock: %d\n", lockValue)
	lockVal := rdb.SetNX(ctx, lock, lockValue, timeout).Val()
	fmt.Printf("acquire lock Val: %v\n", lockVal)
	return lockVal
}

// init 初始化全局redis客户端
func getRdb() *redis.Client {
	if rdb == nil {
		// get redis client directly
		rdb = redis.NewClient(
			&redis.Options{
				Addr:     "localhost:6379",
				Password: "",
				DB:       0,
				// Default is 3.
				Protocol: 3,
			})
	}
	return rdb
}

// task 模拟任务
func task() {
	ctx := context.Background()
	locked := Lock(ctx)
	for !locked {
		fmt.Println("lock failed, try after 1s")
		select {
		// 创建一个持续运行的计时器，即使在循环结束后，计时器仍然会运行，可能导致资源泄露
		// case <-time.Tick(time.Millisecond * 500):
		// 每次循环时都会创建一个新的计时器，使用完后可以被垃圾回收，可以避免资源泄露，推荐
		case <-time.After(time.Millisecond * 500):
			locked = Lock(ctx)
		case <-ctx.Done():
			return
		}
	}
	// 诡异的问题，这样写无效，Lock 没有被重复执行
	// for locked := Lock(ctx); !locked; {
	// 	fmt.Println("lock failed, try after 1s")
	// 	time.Sleep(time.Millisecond * 500)
	// }
	fmt.Println("lock success start task")

	defer func() {
		fmt.Println("Unlock after job finished")
		Unlock(ctx)
	}()

	fmt.Printf("job started: %d\n", count)
	for i := 0; i <= 100; i++ {
		count += i
		time.Sleep(time.Millisecond * 100)
	}
	fmt.Printf("job finished: %d\n", count)
}

func lockTest() {

	goNum := 3
	// 模拟多个任务操作同一个count
	var wg = sync.WaitGroup{}
	wg.Add(goNum)
	for i := 0; i < goNum; i++ {
		go func() {
			defer func() {
				fmt.Println("lockTest exit")
				wg.Done()
			}()
			task()
		}()
	}

	wg.Wait()
	fmt.Printf("lockTest finished: %d\n", count)
}
