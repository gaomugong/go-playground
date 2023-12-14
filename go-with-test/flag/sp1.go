package main

import (
	"flag"
	"fmt"
)

var (
	iflag *int
	bflag *bool
	sflag *string
)

func init() {
	iflag = flag.Int("iflag", 0, "int flag value")
	bflag = flag.Bool("bflag", false, "bool flag value")
	sflag = flag.String("sflag", "default", "string flag value")
}

func main() {
	flag.Parse()

	fmt.Println("int flag:", *iflag)
	fmt.Println("bool flag:", *bflag)
	fmt.Println("string flag:", *sflag)
}
