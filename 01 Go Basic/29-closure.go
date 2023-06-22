package main

import "fmt"

func main() {
	count := 0
	name := "mca"

	increment := func() {
		name := "anwar"
		count++
		fmt.Println(count)
		fmt.Println(name)
	}

	increment()
	fmt.Println(name)
}

// go bulid hello-world.go
// go run hello-world.go
