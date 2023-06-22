package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) married() {
	man.Name = "mr." + man.Name
	fmt.Println(man.Name)

}

func main() {
	mca := Man{"mca"}
	mca.married()
	fmt.Println(mca.Name)
}

// go bulid hello-world.go
// go run hello-world.go
