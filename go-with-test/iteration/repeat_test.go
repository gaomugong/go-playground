package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q, but got %q", expected, repeated)
	}
}

// 基准测试，代码会运行 b.N 次，并测量需要多长时间
// go test -bench=.
// -------------------------------------------------------------------------
// BenchmarkRepeat-10      14509762（测试次数）   82.70 ns/op（单次平均耗时）
// PASS
// ok      iteration       2.657s

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	a := Repeat("a", 5)
	b := Repeat("b", 3)
	fmt.Println(a)
	fmt.Println(b)
	// Output:
	// aaaaa
	// bbb
}
