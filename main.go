package main

import (
	"encoding/json"
	"errors"
	"os"
)

type Greeting string

type Customer struct {
	name string
	sayName func()
	greet func(greeting Greeting)
}

func main() {
	heroes := make(map[string]string)

	heroes["superman"] = "Clark Kent"
	heroes["batman"] = "Bruce Wayne"

	heroes2 := map[string]string{
		"superman": "Clark Kent",
		"batman": "Bruce Wayne",
	}

	var bytes, _ = json.Marshal(heroes2)

	file, err := os.Create("data.txt")
	if err != nil {
		panic(errors.New("failed"))
	}
	defer file.Close()

	file.Write(bytes)
}