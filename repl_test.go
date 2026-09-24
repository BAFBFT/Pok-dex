package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected []string
	}{
		"hello":       {input: "  hello world  ", expected: []string{"hello", "world"}},
		"poke":        {input: "Charmander Bulbasaur PIKACHU ", expected: []string{"charmander", "bulbasaur", "pikachu"}},
		"single_word": {input: "hi", expected: []string{"hi"}},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := cleanInput(tc.input)
			if !reflect.DeepEqual(tc.expected, got) {
				t.Errorf("expected: %v, got: %v", tc.expected, got) // t.Errof logs instead of hard stopping the test
			}
		})
	}
}
