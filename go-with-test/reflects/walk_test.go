package reflects

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	t.Run("v1", func(t *testing.T) {
		expected := "Chris"
		var got []string
		x := struct {
			Name string
		}{expected}

		walk(x, func(input string) {
			got = append(got, input)
		})

		if len(got) != 1 {
			t.Errorf("Expected 1 element, got %d", len(got))
		}

		if got[0] != expected {
			t.Errorf("Expected %s, got %s", expected, got[0])
		}
	})

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name: "Walking one fields struct",
			Input: struct {
				Name string
			}{
				Name: "Chris",
			},
			ExpectedCalls: []string{
				"Chris",
			},
		},
		{
			Name: "Walking two fields struct",
			Input: struct {
				Name string
				City string
			}{
				"Chris", "London",
			},
			ExpectedCalls: []string{
				"Chris", "London",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			var got []string
			walk(c.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, c.ExpectedCalls) {
				t.Errorf("Expected %v, got %v", c.ExpectedCalls, got)
			}
		})
	}

}
