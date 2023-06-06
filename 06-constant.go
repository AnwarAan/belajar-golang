package main

import "fmt"

func main() {
	const firstName string = "mca"
	const age = 12

	fmt.Println(firstName)
	fmt.Println(age)

	const (
		name = "string"
		id   = 001
	)

	fmt.Println(name)
	fmt.Println(id)
}

// go bulid hello-world.go
// go run hello-world.go
