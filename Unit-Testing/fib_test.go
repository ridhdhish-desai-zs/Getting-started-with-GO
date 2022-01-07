package testing

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	testFib := []struct {
		desc     string
		input    int
		expected int
	}{
		{"Case 1", 1, 1},
		{"Case 2", 2, 1},
	}

	for _, tf := range testFib {
		t.Run(tf.desc, func(t *testing.T) {
			output := fib(tf.input)
			if output != tf.expected {
				fmt.Printf("Expected: %d, Got: %d\n", tf.expected, output)
			}
		})
	}
}
