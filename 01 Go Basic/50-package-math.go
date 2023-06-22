package main

import (
	"fmt"
	"math"
)

func main() {
	// round := math.Round()
	fmt.Println(math.Round(1.7))
	fmt.Println(math.Floor(1.7))
	fmt.Println(math.Ceil(1.7))
	fmt.Println(math.Max(17, 30))
	fmt.Println(math.Min(7, 30))

}

// go bulid hello-world.go
// go run hello-world.go
