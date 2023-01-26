package main

import "fmt"

func sort(s []int) {
	slice := s

	// Loop through all items in the slice
	for i := 1; i < len(slice); i++ {
		// If the value to the left is already smaller,
		// do nothing.
		if slice[i] > slice[i-1] {
			continue
		}

		for j := i; j > 0; j-- {
			if slice[j] < slice[j-1] {
				slice[j], slice[j-1] = slice[j-1], slice[j]
				continue
			}
			break
		}
	}
}

func Something(slice []int) {
	copy(slice, []int{123})
}

func main() {
	arr := [100]int{1, 2, 3}
	slice := arr[:]

	Something(slice)

	fmt.Println(slice)
}
