package main

import (
	"bytes"
	"io"
)

type Happy interface {
	io.Reader
	foo() string
}

type happy struct {
	*bytes.Reader
	name string
}

func (h *happy) foo() string {
	return ""
}

type A interface {
	foo()
}

type C struct {
	A
	B
}

type B struct {
	A
}

func (b B) foo() {

}

func main() {
	x := B{}

	x.A.foo()

	z := C{}
	z.A.foo()
	z.B.foo()
}
