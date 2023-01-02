package x

import (
	"fmt"
	"sync"
)

type Account struct {
	balance float32
	lock    sync.Mutex
}

func (account *Account) viewBalance() {
	account.lock.Lock()
	defer account.lock.Unlock()
	fmt.Println(account.balance)
}

func main() {
	account := Account{balance: 21.74}
	group := sync.WaitGroup{}

	// Make 5 thousand parallel requests to view the account's balance
	for i := 1; i <= 5e3; i++ {
		group.Add(1)

		go func() {
			defer group.Done()
			account.viewBalance()
		}()
	}

	group.Wait()
}
