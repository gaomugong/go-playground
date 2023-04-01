package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops uint64

	// Add方法用于向WaitGroup中添加正在等待的协程数量，即向WaitGroup中的计数器添加指定的值。
	// 在这个例子中， wg.Add(1) 向WaitGroup的计数器中添加了1，表示有1个协程正在等待。
	// Done方法用于向WaitGroup中标记一个协程已经完成，即将WaitGroup计数器减1。
	// Wait方法用于等待所有的协程完成。当WaitGroup中的计数器不为0时，Wait方法将会一直等待，直到所有的协程完成。
	var wg sync.WaitGroup

	// start 50 goroutines that increment the counter 1000 times in each
	for i := 0; i < 50; i++ {
		// 计数器加一
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			// 计数器减一
			wg.Done()
		}()
	}

	// wait until all goroutines are done
	wg.Wait()

	// It's safe to access `ops` now because we know
	// no other goroutine is writing to it. Reading
	// atomics safely while they are being updated is
	// also possible, using functions like
	// `atomic.LoadUint64`.
	fmt.Println("ops: ", ops)
}
