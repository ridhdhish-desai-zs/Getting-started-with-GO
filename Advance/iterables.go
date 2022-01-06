package main

import "fmt"

func appendByte(p []byte, data ...byte) []byte {
	println("data: ", cap(p))

	n := len(p)
	m := len(data)
	total := n + m

	slice := make([]byte, (n+m)*2)
	copy(slice, p)
	p = slice

	p = p[0:total]
	copy(p[n:total], data)

	return p
}

func main() {
	fmt.Println("Welcome to the ITERABLES...")

	nums := [5]int{1, 2, 3}

	fmt.Println(nums[0:2])

	// // String: Use slice operator(:) to get the character and use index to get the ASCII value of the character
	// var str string = "Naruto"
	// fmt.Println(str[0:2])
	// fmt.Println(str[0])

	// // Slice operator returns new array with the same address of the source array
	// newNums := nums[0:1]
	// fmt.Println(newNums)

	// fmt.Printf("nums add: %x\n", &nums[0])
	// fmt.Printf("newNums add: %x\n", &newNums[0])

	// newNums[0] = 10
	// fmt.Println(nums, newNums)

	// // We cannot access address of a single character at perticular position in string
	// str2 := str[0:3]
	// fmt.Printf("Str1 add: %x\n", &str)
	// fmt.Printf("str2 add: %x\n", &str2)

	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(s)
	fmt.Printf("T: %T\n", s)
	fmt.Println("Capacity: ", cap(s))
	fmt.Println("Length: ", len(s))

	a := s[1:4]
	fmt.Println("A: ", a[:5])

	s = s[1:4]
	fmt.Printf("T: %T\n", s)
	fmt.Println("Length: ", len(s))
	fmt.Println("Capacity: ", cap(s))
	fmt.Println(s)
	fmt.Println(s[:5])

	fmt.Println(appendByte([]byte{1, 2}, 10, 20, 30, 40))

	var slice []int

	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 4, 5, 6, 7, 8)

	fmt.Println(len(slice), cap(slice))

}
