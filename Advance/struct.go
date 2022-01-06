package main

import "fmt"

type Anime struct {
	name        string
	protagonist string
}

func main() {
	var an Anime

	an.name = "Naruto"
	an.protagonist = "naruto"

	fmt.Println(an)

	var p *Anime
	p = &an
	fmt.Printf("Struct Address: %x", p)
	fmt.Println(", Value: ", *p)

	var ds Anime
	fmt.Println(ds)
}
