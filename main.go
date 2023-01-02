package main

import (
	"fmt"
)

type AnimalName string

type Animal interface {
	MakeSound()
}

type Dog struct {
	name AnimalName
}

func (dog Dog) SayName() {
	fmt.Println(dog.name)
}
func (dog Dog) MakeSound() {
	fmt.Println("Woof!")
}

type Cat struct {
	name AnimalName
}

func(cat Cat) Meow() {
	fmt.Printf("%v says \"Meow!\"", cat.name)
}

func main() {
	// Cat doesn't implement the Animal interface, so
	// Martha's type can't be set to Animal
	martha := Cat{"Martha"}

	martha.Meow()

	// Dog implements the Animal interface, so John can
	// be considered an Animal
	var john Animal = Dog{"John"};

	john.MakeSound()
	john.(Dog).SayName()
}