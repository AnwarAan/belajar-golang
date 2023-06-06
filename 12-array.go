package main

import "fmt"

func main() {

	var names [3]string
	names[0] = "muchamad"
	names[1] = "choirul"
	names[2] = "anwar"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	values := [3]int{
		10, 20, 30,
	}
	fmt.Println(values)
	fmt.Println(len(values))
}

// go bulid hello-world.go
// go run hello-world.go
