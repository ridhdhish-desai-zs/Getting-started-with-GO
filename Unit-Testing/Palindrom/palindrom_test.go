package testing

import (
	"fmt"
	"testing"
)

func TestPalindrom(t *testing.T) {

	testPalindrom := []struct {
		desc     string
		input    string
		expected bool
	}{
		{desc: "Case1", input: "A man, a plan, a canal. Panama", expected: true},
	}

	for _, tp := range testPalindrom {
		t.Run(tp.desc, func(t *testing.T) {
			output := isPalindrom(tp.input)
			fmt.Println(output)
			if output != tp.expected {
				t.Errorf("%t", output)
			}
		})
	}
}
