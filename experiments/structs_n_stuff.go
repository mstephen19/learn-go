package main

import "fmt"

type I1 interface {
	Foo()
}

type I2 interface {
	I1
	Bar()
}

type S1 struct {
	name string
}

func (s *S1) Foo() {
	fmt.Println("Foo")
}

func (s *S1) Bar() {
	fmt.Println("Bar")
}

func main() {
	var x I2 = &S1{}
	s1 := x.(*S1)

	s1.Foo()
	s1.Bar()
}
