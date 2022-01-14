package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func f(v *int, wg *sync.WaitGroup) {
	mutex.Lock()
	*v++
	mutex.Unlock()
	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	var v int = 0
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go f(&v, &wg)
	}
	wg.Wait()
	fmt.Println("Finished", v)
}
