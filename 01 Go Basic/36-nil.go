package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	}
	return map[string]string{
		"name": name,
	}
}

func main() {
	var person map[string]string = nil
	human := newMap("mca")
	fmt.Println(person)
	fmt.Println(human)

	data := newMap("")
	if data == nil {
		fmt.Println("404")
	} else {
		fmt.Println(data)
	}
}

// go bulid hello-world.go
// go run hello-world.go
