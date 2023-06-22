package main

import "fmt"

type Name interface {
	getName() string
}

func sayHello(name Name) {
	fmt.Println("Hello", name.getName())
}

type Person struct {
	Name string
}

func (person Person) getName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) getName() string {
	return animal.Name
}
func main() {
	// sayHello("anwar")
	person := Person{
		Name: "mca",
	}
	sayHello(person)

	animal := Animal{"lizard"}
	sayHello(animal)
}

// go bulid hello-world.go
// go run hello-world.go
