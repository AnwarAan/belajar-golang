package main

import (
	"fmt"
)

func random() interface{} {
	return "OK"
	return false
}

func main() {
	result := random()
	fmt.Println(result)

	resultString := result.(string)
	fmt.Println(resultString)

	// resultInt := result.(int)
	// fmt.Println(resultInt)

	res := random()
	switch value := res.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("uknow")
	}

}

// go bulid hello-world.go
// go run hello-world.go
