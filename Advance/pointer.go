package main

import "fmt"

func main() {
	fmt.Println("Pointers in GO!!")

	var val int = 10
	var p *int
	fmt.Println(p)

	p = &val
	fmt.Printf("Value: %d Memory Address: %x\n", *p, p)

	*p += 10
	fmt.Printf("After Manipulating pointer: %d\nOriginal Variable: %d\n", *p, val)

	var dp **int
	dp = &p
	fmt.Printf("Pointer P Value: %d, Address: %x\n", *dp, dp)
}
