package main

import "fmt"

func main() {
	fmt.Println("Welcome to the ITERABLES...")

	nums := [5]int{1, 2, 3}

	fmt.Println(nums)

	// String: Use slice operator(:) to get the character and use index to get the ASCII value of the character
	var str string = "Naruto"
	fmt.Println(str[0:2])
	fmt.Println(str[0])

	// Slice operator returns new array with the same address of the source array
	newNums := nums[0:1]
	fmt.Println(newNums)

	fmt.Printf("nums add: %x\n", &nums[0])
	fmt.Printf("newNums add: %x\n", &newNums[0])

	newNums[0] = 10
	fmt.Println(nums, newNums)

	// We cannot access address of a single character at perticular position in string
	str2 := str[0:3]
	fmt.Printf("Str1 add: %x\n", &str)
	fmt.Printf("str2 add: %x\n", &str2)
}
