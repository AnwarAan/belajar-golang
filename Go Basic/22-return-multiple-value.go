package main

import "fmt"

func getFullName() (string, string) {
	return "hello", "mca"
}

func skipParameter() (int, int) {
	return 100, 99
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
	bestScore, _ := skipParameter()
	fmt.Println(bestScore)
}

// go bulid hello-world.go
// go run hello-world.go
