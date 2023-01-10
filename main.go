package main

import "fmt"

func main() {
	array := [20]byte{}
	slice := array[:10]
	copy(slice, []byte("foo barrrr"[:3]))
	fmt.Println(array)
}
