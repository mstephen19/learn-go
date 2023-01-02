package main

import (
	"fmt"
	"strconv"
	lib "github.com/mstephen19/start-go/lib"
)

func main() {
	mySlice := []int{1, 2, 3, 4, 5}

	handler := func (acc *string, item int, index int) {
		*acc += strconv.Itoa(item)
	}

	result := lib.Reduce(mySlice, "", handler)

	// Expect the string "12345"
	fmt.Println(*result)

	// Expect 120
	fmt.Println(lib.Factorial(5))

	combinations, _ := lib.NumberOfCombinations(len(mySlice), 2)

	fmt.Println(combinations)
}