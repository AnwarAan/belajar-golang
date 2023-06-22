package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

func countMe() {
	fmt.Println("count me: ")
}

func main() {
	sayHello()

	for i := 0; i < 10; i++ {
		countMe()
	}
}

// go bulid hello-world.go
// go run hello-world.go
