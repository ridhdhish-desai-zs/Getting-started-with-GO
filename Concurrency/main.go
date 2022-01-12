package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

func PrintMe(wg *sync.WaitGroup) {
	mutex.Lock()
	for i := 0; i < 5; i++ {
		fmt.Println(i + 1)
		time.Sleep(time.Second)
	}
	if wg != nil {
		wg.Done()
	}
	mutex.Unlock()
}

func main() {
	fmt.Println("Welcome to the world of asynchronous GO.")

	var wg sync.WaitGroup

	t1 := time.Now()

	wg.Add(1)
	go PrintMe(&wg)
	PrintMe(nil)
	wg.Wait()

	// time.Sleep(time.Second * 2)

	fmt.Println(time.Since(t1))
}
