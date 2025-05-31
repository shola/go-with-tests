package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct{
		name string
		input interface{}
		expectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				name string
				city string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non string field",
			struct {
				name string
				age int
			}{"Chris", 33},
			[]string{"Chris"},
		},
	}
	
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			var got []string
			walk(test.input, func(input string) {
				got = append(got, input)
			})
		
			if !reflect.DeepEqual(got, test.expectedCalls) {
				t.Errorf("got %q, want %q", got, test.expectedCalls)
			}
		})
	}

}