package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

type A interface {
	Foo()
}

type ImplementsA struct {
	Testing string
}

func (x *ImplementsA) Foo() {

}

type EmbedsA struct {
	A
	Name string
}

func main() {
	var x = &EmbedsA{
		A: &ImplementsA{
			Testing: "test",
		},
	}

	// This works. It must be casted
	b := x.A.(*ImplementsA).Testing
}
