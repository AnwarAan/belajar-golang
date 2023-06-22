package main

import "fmt"

func sumAll(numbers ...int) int {
	count := 0
	for _, number := range numbers {
		count += number
	}
	return count
}

func main() {
	fmt.Println(sumAll(10, 20, 30))
	total := sumAll(100, 200)
	fmt.Println(total)

	number := []int{1000, 2000, 3000}
	total = sumAll(number...)
	fmt.Println(total)
}

// go bulid hello-world.go
// go run hello-world.go
