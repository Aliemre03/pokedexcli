package main

import (
	"github.com/Aliemre03/pokedexcli/command"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "Hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "HELLO world",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, cs := range cases {
		actual := command.cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("The lenghts are not equal: %v vs %v", len(actual), len(cs.expected))
			continue
		}
		for i := range actual {
			if actual[i] != cs.expected[i] {
				t.Errorf("The values are not equal: %v vs %v", actual[i], cs.expected[i])
			}
		}
	}

}
