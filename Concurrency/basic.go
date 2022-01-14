package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func PrintMe(i int, wg *sync.WaitGroup, val *int) {
	mutex.Lock()
	*val++
	mutex.Unlock()
	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	val := 0

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go PrintMe(i+1, &wg, &val)
	}

	wg.Wait()

	fmt.Println("Hello At the end: ", val)
}
