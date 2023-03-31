package main

import (
	"fmt"
	"math"
)

type rect struct {
	width, height float64
}

// rect-实现接口
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return (r.width + r.height) * 2
}

type circle struct {
	radius float64
}

// circle-实现接口
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 定义接口geometry
type geometry interface {
	area() float64
	perim() float64
}

// 测试
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
