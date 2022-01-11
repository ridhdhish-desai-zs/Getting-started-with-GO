package main

import "fmt"

type I interface {
	Hello()
}

type MyInt int

type orange struct {
	name string
}

func (m MyInt) Hello() {
	fmt.Println("Hello")
}

func (o orange) Hello() {
	fmt.Println(o)
}

func main() {

	var i I
	// i.Hello()

	var m = MyInt(1)

	i = m
	i.Hello()

	// var o = orange{name: "Melons"}
	// i = &o
	// i.Hello()

	var o1 orange
	i = o1
	o1.Hello()

	// Empty Interface works like any type in TypeScript and we change its type as well
	var b interface{} = 1
	fmt.Println(b)

	b = "Hello"
	fmt.Println(b)

}
