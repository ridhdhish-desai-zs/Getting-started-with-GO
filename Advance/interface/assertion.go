package main

import (
	"fmt"
)

func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v)
	case string:
		fmt.Println(v)
	default:
		fmt.Println(nil)
	}
}

func main() {
	var i interface{} = "Hello"

	s := i.(string)
	fmt.Println(s)

	b, ok := i.(bool)
	fmt.Println(b, ok)

	checkType(10)
	checkType("Hey")
}
