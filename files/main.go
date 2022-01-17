package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to file in golang!!")

	file, err := os.Create("./animes.txt")
	checkNilErr(err)
	defer file.Close()

	dataLength, err := io.WriteString(file, "Naruto is better than sasuke")
	checkNilErr(err)
	fmt.Println("File data length is: ", dataLength)

	dataBytes, err := ioutil.ReadFile("./animes.txt")
	checkNilErr(err)
	fmt.Println("File data is: ", string(dataBytes))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
