package main

import "fmt"

func main() {
	counter := 0

	for counter < 10 {
		counter++
		fmt.Println("loop", counter)
	}

	//init statment
	//post statment
	for c := 1; c < 100; c++ {
		fmt.Println("loop", c)
	}

	slices := []string{"a", "b", "c", "d"}
	for i := 0; i < len(slices); i++ {
		fmt.Println(slices[i])
	}

	//for range
	for index, slice := range slices {
		fmt.Println("index", index, "=", slice)
	}

	for _, slice := range slices {
		fmt.Println(slice)
	}

	person := make(map[string]string)
	person["name"] = "mca"
	person["zipcode"] = "17134"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}

// go bulid hello-world.go
// go run hello-world.go
