package main

import (
	"fmt"
	"sync"
	"time"
)

var mapMutex sync.Mutex

// Create Map Type
type MyMap map[string]int

func (m *MyMap) Add(key string, value int, wg *sync.WaitGroup) {
	mapMutex.Lock()
	time.Sleep(time.Second)
	(*m)[key] = value
	fmt.Println("Added: ", *m)
	mapMutex.Unlock()
	wg.Done()
}

func (m *MyMap) Get(key string, wg *sync.WaitGroup) int {
	mapMutex.Lock()
	time.Sleep(time.Second)
	val := (*m)[key]
	mapMutex.Unlock()
	wg.Done()
	return val
}

func (m *MyMap) Delete(key string, wg *sync.WaitGroup) {
	mapMutex.Lock()
	time.Sleep(time.Second)
	if (*m)[key] == 0 {
		fmt.Println("Cannot delete what's not existed")
	}
	delete(*m, key)
	fmt.Println("Deleted: ", *m)
	mapMutex.Unlock()
	wg.Done()
}

func main() {
	fmt.Println("Awesome Map and it's Methods")

	var wg sync.WaitGroup

	var myMap = make(MyMap)

	wg.Add(3)
	go myMap.Delete("a", &wg)
	go func() {
		val := myMap.Get("a", &wg)
		fmt.Println(val)
	}()
	go myMap.Add("a", 1, &wg)
	wg.Wait()

}
