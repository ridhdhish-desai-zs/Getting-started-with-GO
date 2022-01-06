package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {

	fmt.Println(strings.Fields(s))

	wordCounts := make(map[string]int)

	for _, v := range strings.Fields(s) {
		wordCounts[v] += 1
	}

	return wordCounts
}

func main() {
	fmt.Println(WordCount("Hello x"))
}
