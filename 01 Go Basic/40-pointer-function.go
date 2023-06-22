package main

import "fmt"

type Address struct {
	City, Country string
	Zipcode       int
}

func changeToIndonesia(address *Address) {
	address.Zipcode = 2222

}

func main() {
	direction := Address{City: "bekasi", Country: "malay", Zipcode: 111}
	var goal = &direction
	changeToIndonesia(&direction)
	changeToIndonesia(goal)
	fmt.Println(direction)
}

// go bulid hello-world.go
// go run hello-world.go
