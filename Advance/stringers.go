package main

import "fmt"

type MyAnime struct {
	name     string
	episodes int
}

func (a MyAnime) String() string {
	return fmt.Sprintf("%v has %v Episodes", a.name, a.episodes)
}

func main() {
	a1 := MyAnime{name: "Naruto", episodes: 720}
	a2 := MyAnime{name: "Demon Slayer", episodes: 35}

	fmt.Println(a1)
	fmt.Println(a2)
}
