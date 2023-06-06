package main

import "fmt"

func main() {

	type Name = string
	type Married = bool

	var firstName Name = "muchamad"
	var isMarried = false

	fmt.Println(firstName)
	fmt.Println(Name("choirul"))
	fmt.Println(isMarried)

}

// go bulid hello-world.go
// go run hello-world.go
