package main

import (
	"reflect"
	"testing"
)


func TestWalk(t *testing.T) {
	type Profile struct {
		Age int
		City string
	}

	type Person struct {
		Name string
		Profile Profile
	}

	cases := []struct{
		Name string
		Input interface{}
		expectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"nested fields",
			Person{
				"Chris", 
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
	}
	
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
		
			if !reflect.DeepEqual(got, test.expectedCalls) {
				t.Errorf("got %q, want %q", got, test.expectedCalls)
			}
		})
	}

}