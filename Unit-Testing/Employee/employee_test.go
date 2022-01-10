package employee

import (
	"testing"
)

func TestEmployee(t *testing.T) {
	tp := []struct {
		desc     string
		input    emp
		expected bool
	}{
		{desc: "Case1", input: emp{id: 1, name: "Naruto", hasPan: true, age: 28}, expected: false},
		{desc: "Case2", input: emp{id: 2, name: "Kakashi", hasPan: false, age: 32}, expected: true},
	}

	for _, v := range tp {
		if v.input.name == "Naruto" {
			v.input.setName("Naruto Uzumaki")
			v.input.setAge(16)
			v.input.setHasPan(false)
		}
		t.Run(v.desc, func(t *testing.T) {
			output, _ := checkAge(v.input)

			if output != v.expected {
				t.Errorf("Expected: %v, Got: %v", v.expected, output)
			}
		})
	}

}
