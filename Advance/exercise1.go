package main

import "fmt"

func Square(x int) int {
	return x * x
}

func myMap(slice []int, sq func(x int) int) []int {

	results := make([]int, len(slice))

	for k, v := range slice {
		results[k] = sq(v)
	}

	return results
}

func main() {
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(myMap(slice, Square))
}
