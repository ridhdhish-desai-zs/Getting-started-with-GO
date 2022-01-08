package pm

import (
	"testing"
)

// Test for Fetch all products
func TestGetProducts(t *testing.T) {

	var op1 = []Product{
		{id: 1, name: "Logitech G102 Mouse"},
		{id: 2, name: "New Monitor"},
	}

	tp := []struct {
		desc     string
		expected []Product
	}{
		{desc: "Case1", expected: []Product{}},
		{desc: "Case2", expected: op1},
	}

	for _, v := range tp {
		t.Run(v.desc, func(t *testing.T) {
			output, err := GetProducts()

			if err != nil {
				t.Fatalf("%s", err)
			}

			if len(output) != len(v.expected) {
				t.Errorf("Expected Products: %d, Got Products: %d", len(v.expected), len(output))
			} else {
				for k := range v.expected {
					o1 := v.expected[k]
					e1 := output[k]

					if o1 != e1 {
						t.Errorf("Expected: %d, %s; Got: %d, %s", e1.id, e1.name, o1.id, o1.name)
					}
				}
			}
		})
	}
}

// Test for Add a Product
func TestAddProduct(t *testing.T) {
	tp := []struct {
		desc     string
		id       int
		name     string
		expected Product
	}{
		{desc: "Case1", id: 1, name: "Logitech G102 Mouse", expected: Product{id: 1, name: "Logitech G102 Mouse"}},
		{desc: "Case2", id: 2, name: "New Monitor", expected: Product{id: 2, name: "New Monitor"}},
	}

	for _, v := range tp {
		t.Run(v.desc, func(t *testing.T) {
			output, err := AddProduct(v.id, v.name)

			if err != nil {
				t.Fatalf("%s", err)
			}

			if output != v.expected {
				t.Errorf("Expected: %d, %s; Got: %d, %s", v.expected.id, v.expected.name, output.id, output.name)
			}
		})
	}
}

// Test for Delete a Product
func TestDeleteProduct(t *testing.T) {
	tp := []struct {
		desc     string
		id       int
		expected Product
	}{
		{desc: "Case1", id: 1, expected: Product{id: 1, name: "Logitech G102 Mouse"}},
		{desc: "Case2", id: 10, expected: Product{}},
	}

	for _, v := range tp {
		t.Run(v.desc, func(t *testing.T) {
			output, err := DeleteProduct(v.id)

			if err != nil {
				t.Fatalf("%s", err)
			}

			if output != v.expected {
				t.Errorf("Expected: %d, %s; Got: %d, %s", v.expected.id, v.expected.name, output.id, output.name)
			}
		})
	}
}
