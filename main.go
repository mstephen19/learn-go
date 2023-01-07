package main

import "fmt"

type Animal interface {
	MakeSound()
}

type Cat string

func (c Cat) MakeSound() {
	fmt.Println("Meow")
}

func (c *Cat) Change() {
	*c = "foo"
}

func main() {
	var cat Animal
	cat = Cat("foo")

	kitty, ok := cat.(Cat)

	if !ok {
		panic("whoops")
	}

	fmt.Println(kitty)
}
