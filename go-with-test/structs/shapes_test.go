package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	// got := Area(10.0, 6.0)
	// want := 60.0
	//
	// if got != want {
	// 	t.Errorf("got %.2f, want %.2f", got, want)
	// }
	// 通过Shape接口，checkArea函数从具体的形状类型解耦出来了
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := Shape.Area(shape)
		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10, 6}
		checkArea(t, rectangle, 60.0)
		// got := rectangle.Area()
		// want := 60.0
		//
		// if got != want {
		// 	t.Errorf("got %.2f, want %.2f", got, want)
		// }
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{6}
		// got := circle.Area()
		// want := 113.09733552923255
		//
		// if got != want {
		// 	t.Errorf("got %.2f, want %.2f", got, want)
		// }
		// fmt.Println(circle.Area())
		checkArea(t, circle, 113.09733552923255)
	})

	// 表格驱动测试
	// https://github.com/golang/go/wiki/TableDrivenTests
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"test rect", Rectangle{10, 6}, 60.0},
		{"test circle", Circle{6}, 113.09733552923255},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			checkArea(t, tt.shape, tt.want)
		})
	}
}
