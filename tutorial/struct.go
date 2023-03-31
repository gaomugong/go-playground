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

func main() {
	rect1 := Rect{3, 4}
	rect2 := NewRect(3, 4)
	fmt.Println("area:", rect1.area())
	fmt.Println("area:", rect2.area())
}
