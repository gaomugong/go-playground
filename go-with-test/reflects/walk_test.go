package reflects

import (
	"reflect"
	"testing"
)

type Profile struct {
	Age  int
	City string
}

type Person struct {
	Name    string
	Profile Profile
}

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
		{
			Name: "Walking three fields struct",
			Input: struct {
				Name string
				Age  int
				City string
			}{
				"Chris", 20, "London",
			},
			ExpectedCalls: []string{
				"Chris", "London",
			},
		},
		{
			Name: "Walking nested fields",
			Input: Person{
				"Chris", Profile{20, "London"},
			},
			ExpectedCalls: []string{
				"Chris", "London",
			},
		},
		{
			Name: "Walking pointed fields",
			Input: &Person{
				"Chris", Profile{33, "London"},
			},
			ExpectedCalls: []string{
				"Chris", "London",
			},
		},
		{
			Name: "Walking slice fields",
			Input: []Profile{
				Profile{33, "Chris"},
				Profile{23, "London"},
			},
			ExpectedCalls: []string{
				"Chris", "London",
			},
		},
		{
			Name: "Walking slice fields",
			Input: [2]Profile{
				Profile{33, "Hello"},
				Profile{23, "Tianjin"},
			},
			ExpectedCalls: []string{
				"Hello", "Tianjin",
			},
		},
		{
			Name: "Walking slice fields",
			Input: map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			ExpectedCalls: []string{
				"Bar", "Boz",
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
