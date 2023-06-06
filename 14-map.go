package main

import (
	"fmt"
)

func main() {

	person := map[string]string{"name": "mca", "email": "mca@mail.com"}
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["email"])

	person["title"] = "programmer"
	fmt.Println(person["title"])
	fmt.Println(len(person["title"]))

	book := make(map[string]string)
	book["sastra"] = "aku binatang jalang"
	book["filsafat"] = "dunia sophie"
	book["ekonomi"] = "sadar kaya"
	delete(book, "sastra")
	fmt.Println(book)

}

// go bulid hello-world.go
// go run hello-world.go
