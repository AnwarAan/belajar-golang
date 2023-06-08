package main

import "fmt"

func main() {
	var value32 int32 = 1000000000
	var value64 int64 = int64(value32)
	var value8 int8 = int8(value32) //0

	fmt.Println(value32)
	fmt.Println(value64)
	fmt.Println(value8)

	var name = "string"
	var n byte = name[0]
	var str = string(n)
	fmt.Println(str)

}

// go bulid hello-world.go
// go run hello-world.go
