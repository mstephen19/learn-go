package main

import (
	"fmt"
	// "sync"
)

func countToNumber(channel chan string, label string, num int) {
	for i := 1; i <= num; i++ {
		fmt.Println(label, i)
	}

	channel <- "done"
}

func main() {
	channel := make(chan string)

	go countToNumber(channel, "first", 50)
	go countToNumber(channel, "second", 50)

	<-channel
	<-channel
}
