package main

import (
	"bytes"
	"fmt"
	"io"
	"sync"
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
	*B
}

type B struct {
	Name int
}

func (b B) foo() {
	fmt.Println("ran")
}

type Modifier interface {
	Modify()
}

type L struct {
	lock *sync.Mutex
	name string
}

func (l *L) Modify() {
	l.name = "foo"
}

type M struct {
	lock *sync.Mutex
	name string
}

func (m *M) Modify() {
	m.name = "foo"
}

type S interface{}

func main() {
	var B S = "foo"

	x := B.(string)
}
