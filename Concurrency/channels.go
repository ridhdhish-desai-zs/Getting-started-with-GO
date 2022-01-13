package main

import (
	"fmt"
	"sync"
	"time"
)

var chMutex sync.Mutex

func display(c chan int, val int) {
	chMutex.Lock()
	i := val
	fmt.Println("I: ", i)
	time.Sleep(time.Second)
	c <- int(i)
	c <- 29
	fmt.Println("End: ")
	chMutex.Unlock()
}

func main() {
	fmt.Println("WELCOME TO THE CHANNEL METAVERSE!!")

	ch := make(chan int, 3)
	defer close(ch)

	go display(ch, 5)
	go display(ch, 10)

	val := <-ch
	fmt.Println("First: ", val)
	val = <-ch
	fmt.Println("Second: ", val)
	val = <-ch
	fmt.Println("Third: ", val)

	time.Sleep(time.Second)
}
