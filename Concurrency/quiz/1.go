package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func worker(id int, jobs <-chan string, wg *sync.WaitGroup) {
	for j := range jobs {
		j = "test" + j
		time.Sleep(time.Second)
		fmt.Println("out--->", j)
	}
	wg.Done()
}

func main() {

	task := 10

	var wg sync.WaitGroup

	wg.Add(1)

	jobs := make(chan string, task)

	for j := 1; j <= task; j++ {
		jobs <- strconv.Itoa(j)
	}
	close(jobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, &wg)
	}

	wg.Wait()
}
