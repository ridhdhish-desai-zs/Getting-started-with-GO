package main

import (
	"fmt"
	"log"
	"net/http"
)

type Site struct {
	URL string
}

type Result struct {
	StatusCode int
}

func fetch(wId int, sites <-chan Site, results chan<- Result) {
	for site := range sites {

		log.Printf("ID: %d, URL: %s", wId, site.URL)

		res, _ := http.Get(site.URL)
		results <- Result{StatusCode: res.StatusCode}
	}
}

func main() {
	fmt.Println("THREAD IS LIFE SAVER AS WELL AS TIME SAVER")

	sites := make(chan Site, 2)
	results := make(chan Result, 2)

	for i := 1; i <= 2; i++ {
		go fetch(i, sites, results)
	}

	urls := []string{
		"https://google.com",
		"https://binance.com",
		"https://example.com",
	}

	for _, url := range urls {
		sites <- Site{URL: url}
	}
	close(sites)

	for i := 0; i < 3; i++ {
		res := <-results
		fmt.Println("Site Status: ", res.StatusCode)
	}
}
