package main

import (
	"fmt"
	"strconv"
)

type Order struct {
	id          int
	items       []string
	amount      float64
	shippingFee float64
}

type User struct {
	id     int
	name   string
	wallet float64
	orders []int
}

var orders []Order
var users []User

// Signup User
func Signup(id string, name string) User {
	convertedId, _ := strconv.Atoi(id)
	newUser := User{id: convertedId, name: name}

	users = append(users, newUser)

	return newUser
}

// Add money to the wallet
func (u *User) AddToWallet(amount string) {
	convertedAmount, _ := strconv.Atoi(amount)

	u.wallet += float64(convertedAmount)
}

// Create new order
func NewOrder(id string, items []string, amount float64) Order {
	convertedId, _ := strconv.Atoi(id)
	o := Order{id: convertedId, items: items, amount: amount}

	o.setShippingFee()

	orders = append(orders, o)

	return o
}

// Set the Shipping cost "0" if amount > 500 else "15" rs.
func (o *Order) setShippingFee() {
	if o.amount < 500 {
		o.shippingFee = 15.0
	}
}

// Make payment
func Payment(o Order, u *User) bool {
	if u.wallet < o.amount {
		fmt.Println("Insufficient Fund. Add more fund to your wallet.")
		return false
	}

	u.orders = append(u.orders, o.id)
	return true
}

func main() {
	fmt.Println("Order Management System")

}
