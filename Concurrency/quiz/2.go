package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func f(v *int, wg *sync.WaitGroup) {
	*v++
	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	var v int = 0
	for i := 0; i < 100000; i++ {
		go f(&v, &wg)
	}
	wg.Wait()
	fmt.Println("Finished", v)
}
