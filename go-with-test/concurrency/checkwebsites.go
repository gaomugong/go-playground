package concurrency

import (
	"fmt"
)

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannels := make(chan result)

	for _, url := range urls {
		// v0
		// results[url] = wc(url)

		// v1
		// 在声明匿名函数时所有可用的变量也可在函数体内使用
		go func() {
			// 变量 url 被重复用于 for 循环的每次迭代 —— 每次都会从 urls 获取新值
			// 每个 goroutine 都是 url 变量的引用 —— 它们没有自己的独立副本
			// 所以他们 都 会写入在迭代结束时的 url —— 最后一个 url
			// fmt.Println(url, &url)
			// waat://furhurterwe.geds 0x14000102230
			// waat://furhurterwe.geds 0x14000102230
			// waat://furhurterwe.geds 0x14000102230
			// results[url] = wc(url)
		}()

		// 让循环速度放慢，可以让goroutine在循环间隙里引用到前面的url
		// 但是这样就没有使用goroutine的意义了
		// time.Sleep(time.Second * 2)
		// http://google.com 0x1400008e230
		// http://blog.gypsydave5.com 0x1400008e230
		// waat://furhurterwe.geds 0x1400008e230

		// v2
		go func(url string) {
			// 这里的url是外层url的副本，不会被迭代重写
			fmt.Println(url, &url)
			// waat://furhurterwe.geds 0x1400012e210
			// http://blog.gypsydave5.com 0x14000022060
			// http://google.com 0x1400008c000
			// 偶尔出现：fatal error: concurrent map writes
			// 当我们运行我们的测试时，两个 goroutines 完全同时写入 results map
			// Go 的 Maps 不喜欢多个事物试图一次性写入，所以就导致了 fatal error

			// 这是一种 race condition（竞争条件）借助工具可以进行竞态检测
			// go test -race ./concurrency
			// WARNING: DATA RACE
			// Write at 0x00c000100390 by goroutine 10:
			// concurrency.CheckWebsites.func2()
			// go-with-test/concurrency/checkwebsites.go:48 +0x134
			// results[url] = wc(url)
			resultChannels <- result{url, wc(url)}
		}(url)

	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannels
		results[result.string] = result.bool
	}

	// time.Sleep(time.Second * 2)
	// checkwebsites_test.go:32: Want map[http://blog.gypsydave5.com:false http://google.com:false waat://furhurterwe.geds:true], got map[waat://furhurterwe.geds:true]
	// checkwebsites_test.go:22: Wanted 3, got 1
	return results
}
