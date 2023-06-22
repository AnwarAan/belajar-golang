package main

import (
	"errors"
	"fmt"
)

func divided(value int, div int) (int, error) {
	if div == 0 {
		return div, errors.New("cant't div by zero")
	} else {
		return value / div, nil
	}
}

func main() {
	result, err := divided(10, 0)
	if err == nil {
		fmt.Println("result", result)
	} else {
		fmt.Println("error", err.Error())
	}
}

// go bulid hello-world.go
// go run hello-world.go
