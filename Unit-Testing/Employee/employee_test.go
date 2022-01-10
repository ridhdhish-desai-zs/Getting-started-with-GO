package employee

import (
	"testing"
)

func TestCheckEmployeeAge(t *testing.T) {
	tp := []struct {
		desc     string
		input    emp
		expected bool
	}{
		{desc: "Case1", input: emp{id: 1, name: "Naruto", hasPan: true, age: 17}, expected: false},
		{desc: "Case2", input: emp{id: 2, name: "Kakashi", hasPan: false, age: 32}, expected: true},
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

func TestGetter(t *testing.T) {
	tp := []struct {
		desc     string
		id       int
		expected emp
	}{
		{desc: "Case1", id: 1, expected: emp{id: 1, name: "Naruto", hasPan: false, age: 17}},
		{desc: "Case2", id: 2, expected: emp{id: 2, name: "Pain", hasPan: true, age: 23}},
	}

	for _, v := range tp {
		t.Run(v.desc, func(t *testing.T) {
			output := getDetails(v.id)

			if output != v.expected {
				t.Errorf("Expected: %v, Got: %v", v.expected, output)
			}
		})
	}
}

func TestSetter(t *testing.T) {
	tp := []struct {
		desc     string
		name     string
		id       int
		hasPan   bool
		age      int
		input    emp
		expected emp
	}{
		{desc: "Case1", input: emp{}, name: "Naruto", id: 100, hasPan: true, age: 30, expected: emp{id: 100, name: "Naruto", hasPan: true, age: 30}},
	}

	for _, v := range tp {
		t.Run(v.desc, func(t *testing.T) {
			v.input.setId(100)
			v.input.setAge(30)
			v.input.setName("Naruto")
			v.input.setHasPan(true)

			if v.input != v.expected {
				t.Errorf("Expected: %v, Got: %v", v.expected, v.input)
			}
		})
	}
}
