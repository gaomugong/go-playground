package main

import (
	"fmt"
	"math"
)

// 定义接口geometry
type geometry interface {
	area() float64
	perim() float64
}

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

// 测试
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// 定义一个可以接收任何参数的函数
func myfunc(iface interface{}) {
	fmt.Printf("type: %T, value: %v\n", iface, iface)
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

	// 空接口变量默认值为nil
	var i interface{}
	fmt.Printf("type: %T, value: %v\n", i, i)

	// 空接口可以承载任意类型的值
	i = 1
	myfunc(i)
	i = false
	myfunc(i)
	i = "hello"
	myfunc(i)

	// 类型断言
	i1 := i.(string)
	fmt.Println("i1: ", i1)
	if i2, ok := i.(string); ok {
		fmt.Printf("i is a string var, %s\n", i2)
	}

	// 定义一个可以接收任何类型的array/slice/map/struct
	any1 := make([]interface{}, 5)
	any1[0] = 1
	any1[1] = "hello"
	any1[2] = []int{1, 2, 3, 4, 5}
	myfunc(any1)

	b := make(map[string]interface{})
	b["name"] = "zhangsan"
	b["age"] = 20
	fmt.Println(b)

	// type-switch
	switch i.(type) {
	case int:
		fmt.Println("i is int")
	case string:
		fmt.Println("i is string")
	default:
		fmt.Println("unknown type")
	}
}
