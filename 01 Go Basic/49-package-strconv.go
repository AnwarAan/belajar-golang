package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	number, err := strconv.ParseInt("1000000", 10, 64)
	value := strconv.FormatInt(10000, 10)
	valueInt, err := strconv.Atoi("200000")

	if err == nil {
		fmt.Println(boolean)
		fmt.Println(number)
		fmt.Println(value)
		fmt.Println(valueInt)
	} else {
		fmt.Println(err)
	}
}

// go bulid hello-world.go
// go run hello-world.go
