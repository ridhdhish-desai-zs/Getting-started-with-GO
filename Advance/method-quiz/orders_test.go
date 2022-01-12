package ordertTest

import (
	"reflect"
	"testing"
)

func TestOrder(t *testing.T) {
	tp := []struct {
		desc     string
		id       int
		items    []string
		amount   float64
		expected Order
	}{
		{desc: "Case1", id: 1, items: []string{"ABC"}, amount: 260, expected: Order{id: 1, items: []string{"ABC"}, amount: 260, shippingFee: 0.0}},
		{desc: "Case2", id: 2, items: []string{"ABC"}, amount: 200, expected: Order{id: 2, items: []string{"ABC"}, amount: 200, shippingFee: 10.0}},
	}

	for _, v := range tp {
		o := *New(v.id, v.items, v.amount)

		t.Run(v.desc, func(t *testing.T) {
			if !reflect.DeepEqual(o, v.expected) {
				t.Errorf("Expected: %v, Got: %v", v.expected, o)
			}
		})
	}
}
