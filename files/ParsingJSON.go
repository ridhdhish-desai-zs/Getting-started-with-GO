package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// FIXME: Returning array of empty objects

type Animes struct {
	Animes []Anime `json: "animes"`
}

type Anime struct {
	Name        string `json: "name"`
	Protagonist string `json: "protagonist"`
	Antagonist  string `json: "antagonist"`
}

func main() {
	fmt.Println("Read the content of JSON file in STRUCT!!")

	var animes Animes

	file, err := os.Open("./anime.json")
	checkNilErr(err)
	defer file.Close()

	dataBytes, err := ioutil.ReadAll(file)
	checkNilErr(err)
	// fmt.Println(string(dataBytes))

	err = json.Unmarshal(dataBytes, &animes)
	checkNilErr(err)

	for _, v := range animes.Animes {
		fmt.Printf("Name: %s, Protagonist: %s, Antagonist: %s\n", v.Name, v.Protagonist, v.Antagonist)
	}

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
