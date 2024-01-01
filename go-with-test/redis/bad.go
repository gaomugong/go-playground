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

// relaseLock 释放分布式锁
// 这里并没有考虑误删除，至少要有对锁值的判定
func relaseLock(ctx context.Context) {
	rdb := getRdb()
	fmt.Printf("release lock: %v\n", rdb.Get(ctx, lock).Val())
	fmt.Println(rdb.Del(ctx, lock))
}

// Lock 获取分布式锁
// 至少需要支持指定key的名字，让使用该组件的人可以创建多把锁
// 这里可以考虑用uuid来生成唯一锁值，减少分布式模式下的value碰撞机会
// 综上，这里需要把面向过程改为面向对象（结构体），并将锁的key和value存储起来，也确保释放操作能够拿到锁值
func acquireLock(ctx context.Context) bool {
	rdb := getRdb()
	lockValue := rand.Int()
	fmt.Printf("acquire lock: %d\n", lockValue)
	lockVal := rdb.SetNX(ctx, lock, lockValue, timeout).Val()
	fmt.Printf("acquire lock Val: %v\n", lockVal)
	return lockVal
}

// init 初始化全局redis客户端
// TODO: 这里的rdb是否需要关闭？
func getRdb() *redis.Client {
	if rdb == nil {
		// get redis client directly
		return redis.NewClient(
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
// 这里一开始只是用了time.Sleep来实现阻塞获取锁，改成select+time.After方式后，能够避免阻塞协程
func task() {
	// ctx := context.Background()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	locked := acquireLock(ctx)
	for !locked {
		fmt.Println("lock failed, try after 200ms")
		select {
		// 创建一个持续运行的计时器，即使在循环结束后，计时器仍然会运行，可能导致资源泄露
		// case <-time.Tick(time.Millisecond * 500):
		// 每次循环时都会创建一个新的计时器，使用完后可以被垃圾回收，可以避免资源泄露，推荐
		case <-time.After(time.Millisecond * 200):
			locked = acquireLock(ctx)
		case <-ctx.Done():
			fmt.Println("task canceled")
			return
		}
	}
	// TODO: 诡异的问题，这样写无效，Lock 没有被重复执行
	// for locked := Lock(ctx); !locked; {
	// 	fmt.Println("lock failed, try after 1s")
	// 	time.Sleep(time.Millisecond * 500)
	// }
	fmt.Println("lock success start task")

	defer func() {
		fmt.Println("relaseLock after job finished")
		relaseLock(ctx)
	}()

	fmt.Printf("job started: %d\n", count)
	for i := 0; i <= 100; i++ {
		count += i
		time.Sleep(time.Millisecond * 10)
	}
	fmt.Printf("job finished: %d\n", count)
}

func lockTest() {

	goNum := 2
	// 模拟多个任务操作同一个count，count最终的值应为5050+5050
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
