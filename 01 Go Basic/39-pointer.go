package main

import "fmt"

type Living struct {
	City, Country string
	Zipcode       int
}

func main() {
	address1 := Living{"bekasi", "indoneisa", 17134}
	address2 := address1
	address3 := &address1

	address2.City = "jakarta"
	address3.Country = "Malay"

	*address3 = Living{"depok", "indonesia", 11111}
	var address4 *Living = new(Living)

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address4)

}

// go bulid hello-world.go
// go run hello-world.go
