package structs

import "math"

func Perimeter(weight float64, height float64) float64 {
	return 2 * (weight + height)
}

// func Area(w float64, h float64) float64 {
// 	return w * h
// }

func Area(r Rectangle) float64 {
	return r.Width * r.Height
}

// ./shapes.go:15:6: Area redeclared in this block
// func Area(c Circle) float64 {
// 	return 3.14 * c.r * c.r
// }

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	r float64
}

// Area shapes_test.go:34:17: circle.Area undefined (type Circle has no field or method Area)
func (c Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

type Shape interface {
	Area() float64
}
