package main

import (
	"fmt"
	"regexp"
)

func main() {
	a := removeAllBsPrecedingA("babxyzabba")
	fmt.Println(a)

	a = removeAllBsPrecedingA("baba")
	fmt.Println(a)
}

func removeAllBsPrecedingA(input string) (out string) {
	re := regexp.MustCompile(`ba`)

	var a []byte

	a = re.ReplaceAll([]byte(input), []byte("a"))

	out = string(a)

	return
}
