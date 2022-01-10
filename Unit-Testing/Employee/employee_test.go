package employee

import "testing"

func TestEmployee(t *testing.T) {
	tp := []struct {
		desc     string
		input    emp
		expected bool
	}{
		{desc: "Case1", input: emp{name: "Naruto", age: 16}, expected: false},
		{desc: "Case2", input: emp{name: "Kakashi", age: 32}, expected: true},
	}

	for _, v := range tp {
		t.Run(v.desc, func(t *testing.T) {
			output, _ := checkAge(v.input)

			if output != v.expected {
				t.Errorf("Expected: %v, Got: %v", v.expected, output)
			}
		})
	}

}
