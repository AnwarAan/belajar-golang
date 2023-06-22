package main

import "fmt"

type Blacklist func(string) bool

func registrtion(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("blocked", name)
	}
	fmt.Println("wlcome", name)
}

func main() {
	blacklist := func(name string) bool {
		return name == "anwar"
	}
	registrtion("anwar", blacklist)
}

// go bulid hello-world.go
// go run hello-world.go
