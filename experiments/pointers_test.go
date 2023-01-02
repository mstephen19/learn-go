package x

import "fmt"

type Person struct {
	name string
}

func handlePerson(person *Person) {
	person.name = "foo"
}

func handleString(str *string) {
	*str = "foo"
}

func z() {
	person := Person{"Mark"}

	handlePerson(&person)

	fmt.Println(person)

	str := "foo"

	handleString(&str)

	fmt.Println(str)
}
