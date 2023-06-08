package main

import "fmt"

func main() {

	initial := "m"
	switch initial {
	case "m":
		fmt.Println("mca")
	case "c":
		fmt.Println("choirul")
	default:
		fmt.Println("anwar")

	}

	switch length := len(initial); length == 0 {
	case false:
		fmt.Println("zero")
	default:
		fmt.Println("undername")
	}

	length := 10
	switch {
	case length < 4:
		fmt.Println("small")
	case length > 4 && length < 7:
		fmt.Println("medium")
	case length > 7:
		fmt.Println("large")
	default:
		fmt.Println("null")

	}

}

// go bulid hello-world.go
// go run hello-world.go
