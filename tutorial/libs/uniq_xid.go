package main

import (
	"log"

	"github.com/rs/xid"
)

// 生成全局唯一字符ID
func main() {
	guid := xid.New()
	log.Println(guid.String())
	log.Println(guid.Machine())
	log.Println(guid.Pid())
	log.Println(guid.Time())
	log.Println(guid.Counter())
}
