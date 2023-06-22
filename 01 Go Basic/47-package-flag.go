package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "set host")
	username := flag.String("username", "root", "set user")

	flag.Parse()
	fmt.Println(*host, *username)
}

// go bulid hello-world.go
// go run hello-world.go

// go run flag.go -h=localhost
