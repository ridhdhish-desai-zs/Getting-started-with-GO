package ordertTest

type Order struct {
	id          int
	items       []string
	shippingFee float64
	amount      float64
}

func New(id int, items []string, amount float64) *Order {
	o := Order{id: id, items: items, amount: amount}
	o.setShippingFee()
	// po := &o

	return &o
}

func (o *Order) setShippingFee() {
	if o.amount < 250 {
		o.shippingFee = 10.0
	} else {
		o.shippingFee = 0.0
	}
}
