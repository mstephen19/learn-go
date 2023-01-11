package main

type A interface {
	Foo()
}

type ImplementsA struct {
}

func (x *ImplementsA) Foo() {

}

type EmbedsA struct {
	A
	test string
}

func main() {
	var x A = EmbedsA{
		A: &ImplementsA{},
	}

	x.Foo()
}
