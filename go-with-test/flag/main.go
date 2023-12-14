package main

import (
	"flag"
	"fmt"
)

var (
	intflag    int
	boolflag   bool
	stringflag string
)

func init() {
	flag.IntVar(&intflag, "intflag", 0, "int flag value")
	flag.BoolVar(&boolflag, "boolflag", false, "bool flag value")
	flag.StringVar(&stringflag, "stringflag", "default", "string flag value")
}

// ./main -intflag 12 -boolflag -stringflag test
// int flag: 12
// bool flag: true
// string flag: test
func main() {
	flag.Parse()

	fmt.Println("int flag:", intflag)
	fmt.Println("bool flag:", boolflag)
	fmt.Println("string flag:", stringflag)
}
