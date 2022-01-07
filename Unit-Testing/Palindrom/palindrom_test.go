package testing

import (
	"testing"
)

func TestPalindrom(t *testing.T) {

	testPalindrom := []struct {
		desc     string
		input    string
		expected bool
	}{
		{desc: "Case1", input: "aba", expected: true},
		{desc: "Case2", input: "", expected: true},
		{desc: "Case3", input: "abc", expected: false},
		{desc: "Case4", input: "a", expected: true},
	}

	for _, tp := range testPalindrom {
		t.Run(tp.desc, func(t *testing.T) {
			output := isPalindrom(tp.input)
			if output != tp.expected {
				t.Errorf("%t", output)
			}
		})
	}
}
