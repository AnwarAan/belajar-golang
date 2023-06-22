package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	luffy := Customer{
		Name: "Luffy",
	}
	luffy.sayHello("Zoro")
}

// go bulid hello-world.go
// go run hello-world.go
