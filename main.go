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

func main() {
	var x Happy = &happy{}

	y := x.(*happy)
}
