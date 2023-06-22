package main

import "fmt"

func main() {

	var a = 10
	var b = 5

	inc := a + b
	a += 11
	b++
	maybe := true

	fmt.Println(inc)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(!maybe)

}

// go bulid hello-world.go
// go run hello-world.go
