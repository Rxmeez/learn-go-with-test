package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct{ Name string }{"Rameez"},
			[]string{"Rameez"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Rameez", "London"},
			[]string{"Rameez", "London"},
		},
		{
			"Stuct with non string field",
			struct {
				Name string
				Age  int
			}{"Rameez", 26},
			[]string{"Rameez"},
		},
		{
			"Nested fields",
			Person{
				"Rameez",
				Profile{26, "London"},
			},
			[]string{"Rameez", "London"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

}
