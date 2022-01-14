package main

import "fmt"

func main() {
	var input interface{}

	type z struct {
		Id   int
		Name string
	}

	user := z{1, "Hey"}
	input = user

	out := input.(z)
	fmt.Println(out)
}
