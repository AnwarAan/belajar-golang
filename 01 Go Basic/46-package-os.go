package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("args")
	fmt.Println(args)
	fmt.Println(args[0])
	fmt.Println(args[1])

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("hostname", name)
	} else {
		fmt.Println("err", err)
	}
}

// go bulid hello-world.go
// go run hello-world.go

// go run os.go mca
// hostname
