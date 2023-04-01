package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("/tmp/defer.txt")
	// close file after write file and before exit
	defer closeFile(f)

	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	if f, err := os.Create(p); err != nil {
		panic(err)
	} else {
		return f
	}
}
func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprint(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	if err := f.Close(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
