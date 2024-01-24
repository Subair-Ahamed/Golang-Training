package main

import (
	"errors"
	"fmt"
	"sync"
)

type Bank struct {
	balance int
	sync.RWMutex
}

func (b *Bank) Balance() int {
	b.RLock()
	defer b.RUnlock()
	return b.balance
}

func (b *Bank) Deposit(amount int) int {
	b.RLock()
	defer b.RUnlock()
	b.balance = b.balance + amount
	return b.balance

}

func (b *Bank) Withdraw(amount int) error {
	b.RLock()
	defer b.RUnlock()
	if amount <= b.balance {
		b.balance = b.balance - amount
		return nil
	}
	return errors.New("Insufficient Balance")
}

func main() {

	acc := Bank{balance: 1000}

	channel := make(chan struct{})

	// mutex := sync.Mutex{}

	go func() {
		// mutex.Lock()
		// defer mutex.Unlock()
		defer func() { channel <- struct{}{} }()
		err := acc.Withdraw(500)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("successfully withdrawn 500")

	}()

	go func() {
		// mutex.Lock()
		// defer mutex.Unlock()
		defer func() { channel <- struct{}{} }()
		err := acc.Withdraw(700)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Successfully withdrawn 700")
	}()

	<-channel
	<-channel

}
