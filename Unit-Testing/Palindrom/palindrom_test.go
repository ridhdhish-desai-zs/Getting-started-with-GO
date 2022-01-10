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
		{desc: "Case1", input: "0A man, a plan, a canal. Panama9", expected: false},
		{desc: "Case2", input: "0A man, a plan, a canal. Panama0", expected: true},
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
