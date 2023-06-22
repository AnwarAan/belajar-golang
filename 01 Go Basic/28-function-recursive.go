package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorial(value int) int {
	if value == 1 {
		return 1
	}
	return value * factorial(value-1)
}

func main() {
	fmt.Println(factorialLoop(5))
	fmt.Println(factorial(5))
}

// go bulid hello-world.go
// go run hello-world.go
