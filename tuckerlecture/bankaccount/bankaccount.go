package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Account struct {
	balance int
	mutex   *sync.Mutex
}

func (a *Account) Widthraw(val int) {
	//a.mutex.Lock()
	a.balance -= val
	//a.mutex.Unlock()
}

func (a *Account) Deposit(val int) {
	//a.mutex.Lock()
	a.balance += val
	//a.mutex.Unlock()
}

func (a *Account) Balance() int {
	return a.balance
}

var accounts []*Account
var globalLock *sync.Mutex

func Transfer(sender, receiver int, money int) {
	globalLock.Lock()
	accounts[sender].Widthraw(money)
	accounts[receiver].Deposit(money)
	globalLock.Unlock()
	fmt.Println("Transfer", sender, receiver, money)
}

func GetTotalBalance() int {
	total := 0
	for i := 0; i < len(accounts); i++ {
		total += accounts[i].Balance()
	}
	return total
}

func RandomTransfer() {
	var sender, balance int
	for {
		sender = rand.Intn(len(accounts))
		balance = accounts[sender].Balance()
		if balance > 0 {
			break
		}
	}

	var receiver int
	for {
		receiver = rand.Intn(len(accounts))
		if sender != receiver {
			break
		}
	}

	money := rand.Intn(balance)
	Transfer(sender, receiver, money)
}

func GoTransfer() {
	for {
		RandomTransfer()
	}
}

func PrintTotalBalance() {
	fmt.Printf("Total: %d\n", GetTotalBalance())
}

func main() {
	globalLock = &sync.Mutex{}
	for i := 0; i < 20; i++ {
		accounts = append(accounts, &Account{balance: 1000, mutex: &sync.Mutex{}})
	}

	go func() {
		for {
			Transfer(0, 1, 100)
		}
	}()

	go func() {
		for {
			Transfer(1, 0, 100)
		}
	}()

	for {
		time.Sleep(100 * time.Millisecond)
	}
	/*
		PrintTotalBalance()

		for i := 0; i < 10; i++ {
			go GoTransfer()
		}
	*/
}
