package main

import "fmt"

func update(words *[]string) {
	slice := *words
	for i, word := range slice {
		slice[i] = fmt.Sprintf("%s-test", word)
	}
}

func main() {
	myWords := []string{"foo", "bar", "baz"}
	update(&myWords)

	fmt.Println(myWords)
}
