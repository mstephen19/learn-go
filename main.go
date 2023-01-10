package main

import "fmt"

func copyStuff(s []byte) {
	// Since our slice has a length of zero,
	// the byte 97 is added to the beginning
	// of the slice's underlying array.
	// But since we are reassigning it, we create
	// a brand new slice. However, this brand
	// new slice still has the same underlying array.
	s = append(s, 'a')
}

func main() {
	array := [20]byte{}
	slice := array[:0]
	copyStuff(slice)
	fmt.Println(array)
	fmt.Println(slice)
}
