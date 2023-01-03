package x

import (
	"fmt"
	"sync"
	"time"
)

func wait(channel *chan string) {
	// Wait for a message to be received on the channel
	msg := <-*channel
	fmt.Println("received", msg)
	// Sleep for 2 seconds
	time.Sleep(2 * time.Second)
	fmt.Println("slept")
	// Send a message onto the channel
	*channel <- "b"
	fmt.Println("sent message")
}

func l() {
	var channel = make(chan string)
	var wg sync.WaitGroup

	go wait(&channel)

	// Send a message on the channel letting it know that it can start
	channel <- "a"

	// Wait for a message on the channel
	fmt.Println(<-channel)

	wg.Add(1)
	go func() {
		fmt.Println(<-channel)
		wg.Done()
	}()

	close(channel)
	wg.Wait()
}
