package main

import "fmt"

func getCompleteName() (fName, mName, lName string) {
	fName = "muchamad"
	mName = "choirul"
	lName = "anwar"
	return
}

func main() {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}

// go bulid hello-world.go
// go run hello-world.go
