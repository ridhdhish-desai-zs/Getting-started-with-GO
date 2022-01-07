package testing

import (
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
		{desc: "Case 3", input: -1},
	}

	for _, tf := range testFib {
		t.Run(tf.desc, func(t *testing.T) {
			output, err := fib(tf.input)

			// Uncomment only after implementation of error case
			if err != nil {
				t.Errorf("Error: %s", err)
			}

			if output != tf.expected {
				t.Errorf("Expected: %d, Got: %d\n", tf.expected, output)
			}
		})
	}
}
