package string

import (
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
