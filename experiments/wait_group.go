package x

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Println("Worker starting", id)
	time.Sleep(5 * time.Second)
	fmt.Println("Worker finished", id)
}

func b() {
	group := sync.WaitGroup{}

	for i := 1; i <= 5; i++ {
		group.Add(1)

		id := i

		go func() {
			defer group.Done()
			worker(id)
		}()
	}

	group.Wait()
}
