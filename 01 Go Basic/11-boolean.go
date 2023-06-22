package main

import "fmt"

func main() {

	absent := 80
	score := 90

	passAbsent := absent >= 75
	passScore := score >= 45

	graduate := passAbsent && passScore

	fmt.Println(graduate)

}

// go bulid hello-world.go
// go run hello-world.go
