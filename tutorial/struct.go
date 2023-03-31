package main

import "fmt"

type Rect struct {
	width, height int
}

func NewRect(width, height int) *Rect {
	return &Rect{width: width, height: height}
}

func (r Rect) area() int {
	return r.width * r.height
}

func (r *Rect) area1() int {
	return r.width * r.height
}

func main() {
	rect1 := Rect{3, 4}
	rect2 := NewRect(3, 4)
	rect3 := &rect1

	fmt.Println("area:", rect1.area())
	fmt.Println("area:", rect2.area())

	// Go automatically handles conversion between values
	// and pointers for method calls.
	fmt.Println("area:", rect1.area1())
	fmt.Println("area:", rect3.area1())
}
