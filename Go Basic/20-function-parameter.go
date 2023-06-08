package main

import "fmt"

func sayGreet(greet string, name string) {
	fmt.Println("good", greet, name)
}

func main() {
	greet := "morning"
	sayGreet(greet, "mca")
}

// go bulid hello-world.go
// go run hello-world.go
