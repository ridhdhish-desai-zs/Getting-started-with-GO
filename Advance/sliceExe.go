package main

import "fmt"

func Pic(dy, dx int) [][]uint8 {
	slice := make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		slice[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			slice[i][j] = uint8(i * j)
		}
	}

	return slice
}

func main() {
	pic := Pic(3, 3)
	fmt.Println(pic)
}
