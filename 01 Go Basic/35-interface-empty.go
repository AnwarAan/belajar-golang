package main

import "fmt"

func input(input string) interface{} {
	return input
}

func main() {
	var data interface{} = input("search")
	fmt.Println(data)
}

// go bulid hello-world.go
// go run hello-world.go
