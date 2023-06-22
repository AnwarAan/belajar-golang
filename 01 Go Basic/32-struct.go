package main

import "fmt"

func main() {
	type Customer struct {
		name, address, zipcode string
		age                    int
		active                 bool
	}

	var data Customer
	data.name = "mca"
	data.age = 22

	fmt.Println(data)
	fmt.Println(data.name)

	user := Customer{
		name: "anwar",
	}
	fmt.Println(user)

	admin := Customer{"mcnwr76", "bekasi", "112", 2, true}
	fmt.Println(admin)
}

// go bulid hello-world.go
// go run hello-world.go
