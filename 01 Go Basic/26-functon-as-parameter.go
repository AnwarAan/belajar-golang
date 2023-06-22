package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
	filterName := filter(name)
	fmt.Print("Hello ", filterName)
}

func spamFilter(name string) string {
	if name == "dog" {
		return "..."
	}
	return name
}

//alisa declaration
type Filter func(string) string

func filteringSayHello(name string, filter Filter) {
	filterName := filter(name)
	fmt.Println("morning mr " + filterName)
}

func main() {
	// fmt.Println(sayHelloWithFilter("dog", spamFilter))
	sayHelloWithFilter("dog", spamFilter)
	filter := spamFilter
	sayHelloWithFilter("mca ", filter)

	filteringSayHello("dog", spamFilter)
}

// go bulid hello-world.go
// go run hello-world.go
