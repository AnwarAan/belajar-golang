package main

import "fmt"

func getGoodBy(name string) string {
	return "good bye " + name
}

func main() {
	goodBy := getGoodBy("mcnwr")
	fmt.Println(goodBy)
}

// go bulid hello-world.go
// go run hello-world.go
