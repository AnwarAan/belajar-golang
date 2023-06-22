package main

import "fmt"

func main() {

	person := map[string]string{"name": "mca", "email": "mca@mail.com"}
	if person["name"] != "mca" {
		fmt.Println("nama depan")
	} else if person["email"] == "mca@mail.com" {
		fmt.Println("email pribadi")
	} else {
		fmt.Println("wrong")
	}

	if length := len(person); length > 5 {
		fmt.Println("lebih dari lima")
	} else {
		fmt.Println("kurang dari lima")
	}

}

// go bulid hello-world.go
// go run hello-world.go
