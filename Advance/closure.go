package main

import "fmt"

func a() func() string {

	name := "Ridhdhish"

	return func() string {
		fmt.Println("Function Returning from a function.")
		return name
	}
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		c := a
		a, b = b, a+b
		return c
	}
}

func main() {
	fmt.Println("Closurrrrrrrrsesss......")

	name := a()()
	fmt.Println(name)

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
