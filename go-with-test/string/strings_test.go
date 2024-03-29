package string

import (
	"fmt"
	"strings"
	"testing"
)

func TestContains(t *testing.T) {
	// t.Run("match", func(t *testing.T) {
	// 	got := strings.Contains("hello world", "world")
	// 	want := true
	//
	// 	if got != want {
	// 		t.Errorf("got %v, want %v", got, want)
	// 	}
	// })

	type args struct {
		s      string
		substr string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"contains", args{"hello world", "world"}, true},
		{"not contains", args{"hello world", "pitou"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := strings.Contains(tt.args.s, tt.args.substr)
			if got != tt.want {
				t.Errorf("%s: want %v, got %v", tt.name, tt.want, got)
			}
		})
	}
}

func stringArrayEqual(arr1, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func TestSplit(t *testing.T) {
	type args struct {
		s   string
		sep string
	}

	// a := strings.Join([]string{"hello", "world"}, ",")
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"split by space", args{"hello world", " "}, []string{"hello", "world"}},
		{"split by comma", args{"hello,world", ","}, []string{"hello", "world"}},
	}

	// [hello world] [hello world]
	//    strings_test.go:65: split by comma: want ["hello" "world"], got ["hello" "world"]
	// --- FAIL: TestSplit (0.00s)
	//    --- FAIL: TestSplit/split_by_comma (0.00s)
	//
	//
	// FAIL
	// 好奇怪的单元测试（reflect.DeepEqual）

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := strings.Split(tt.args.s, tt.args.sep)
			// fmt.Println(got, tt.want)
			// if reflect.DeepEqual(got, tt.want) {
			if !stringArrayEqual(got, tt.want) {
				t.Errorf("%s: want %q, got %q", tt.name, tt.want, got)
			}
		})
	}
}

func TestJoin(t *testing.T) {
	type args struct {
		elems []string
		sep   string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"join by comma", args{[]string{"hello", "world"}, ","}, "hello,world"},
		{"join by space", args{[]string{"hello", "world"}, " "}, "hello world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := strings.Join(tt.args.elems, tt.args.sep)
			if got != tt.want {
				t.Errorf("%s: want %q, got %q", tt.name, tt.want, got)
			}
		})
	}
}

func ExampleStringConverter() {
	fmt.Println(strings.ToLower("GAOMUGONG"))
	fmt.Println(strings.Trim(" GAOMUGONG ", " "))
	fmt.Println(strings.TrimSpace("\t GAOMUGONG \t\n"))
	fmt.Println(strings.ToUpper("pitou"))
	fmt.Println(strings.HasPrefix("gao song", "gao"))
	fmt.Println(strings.HasSuffix("gao song", "song"))
	// 	Output:
	// gaomugong
	// GAOMUGONG
	// GAOMUGONG
	// PITOU
	// true
	// true
}
