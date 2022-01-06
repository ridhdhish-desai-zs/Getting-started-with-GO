package main

import "fmt"

func Square(x int) int {
	return x * x
}

func myMap(slice []int, sq func(x int) int) map[int]int {
	squareMap := make(map[int]int)

	for _, v := range slice {
		squareMap[v] = sq(v)
	}

	return squareMap
}

func main() {
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(myMap(slice, Square))
}
