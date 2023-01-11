package main

import (
	"fmt"
	"sync"
)

type UserAccount interface {
	IncreaseBalance(money float32)
}

type Account struct {
	*sync.Mutex
	balance float32
}

type AccountManager struct {
	UserAccount
}

func (account *Account) IncreaseBalance(money float32) {
	account.Lock()
	defer account.Unlock()

	account.balance += money
}

func (account *Account) DecreaseBalance(money float32) {
	account.Lock()
	defer account.Unlock()

	account.balance -= money
}

func main() {
	var account *AccountManager = &AccountManager{
		&Account{
			Mutex: &sync.Mutex{},
		},
	}

	a := account.UserAccount.(*Account)
	a.balance

	_, ok := account.(*Account)
	if !ok {
		panic("oops!")
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i <= 1000; i++ {
			account.IncreaseBalance(1)
		}
		fmt.Println("done")
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= 1000; i++ {
			account.IncreaseBalance(1)
		}
		fmt.Println("done")
	}()

	wg.Wait()
}
