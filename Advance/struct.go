package main

import (
	"fmt"
	"math"
)

type Anime struct {
	name        string
	protagonist string

	friend *Anime
}

type Location struct {
	lat  float64
	long float64
}

// lat/long distance finding
func FindDistance(l1, l2 Location) float64 {

	latValue := math.Pow(l1.lat-l2.lat, 2)
	longValue := math.Pow(l1.long-l2.long, 2)

	fmt.Println(latValue)

	totalDistance := math.Sqrt(latValue - longValue)

	return totalDistance
}

func main() {
	// We Cannot add struct methods
	// We cannot define struct inside struct; but we can only use a struct type
	// We can use pointers as a property of the struct type

	var an Anime
	var ds Anime

	ds.name = "Demon Slayer"
	ds.protagonist = "Tanjiro"

	an.name = "Naruto"
	an.protagonist = "naruto"
	an.friend = &ds

	fmt.Println(an)

	var p *Anime
	p = &an
	// fmt.Printf("Struct Address: %x", p)
	fmt.Println(", Value: ", *p)
	fmt.Println(an)

	fmt.Println(ds)

	// Get 2 locations
	var surat, bangalore Location

	surat.lat = 53.3
	surat.long = -1.7

	bangalore.lat = 52.2
	bangalore.long = -1.6

	fmt.Printf("Distance is: %.2f kilometers\n", FindDistance(surat, bangalore))
}
