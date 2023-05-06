package main

import (
	"fmt"

	// util 本身也还在自己的本地上开发，并没有上传到 github， 那么 demo 包在调试过程中肯定是无法找到 util 包的
	// 方案一：可以通过在 go.mod 中使用 replace 来重定向
	// 方案二：工作区, go.work
	"github.com/tutor-gowork/util"
)

func main() {
	fmt.Println(util.SayHello("明哥"))
}
