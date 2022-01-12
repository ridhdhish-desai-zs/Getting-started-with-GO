package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

type Account struct {
	accNo  int
	amount int
	holder string
}

func (a *Account) deposit(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	for i := 0; i < 2; i++ {
		time.Sleep(time.Second)
		a.amount += value
		fmt.Println("Deposite: ", a.amount)
	}
	mutex.Unlock()
	wg.Done()
}

func (a *Account) withdraw(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	time.Sleep(time.Second)
	a.amount -= value
	fmt.Println("Withdraw: ", a.amount)
	mutex.Unlock()
	wg.Done()
}

func main() {
	fmt.Println("Bank Transaction System (Reliable)")

	var a1 = Account{accNo: 1, amount: 1000, holder: "Naruto"}
	var wg sync.WaitGroup

	wg.Add(3)
	go a1.withdraw(400, &wg)
	go a1.deposit(200, &wg)
	go a1.withdraw(100, &wg)
	wg.Wait()

	fmt.Printf("After Transaction: %v\n", a1)
}
