package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("anwar")
	data.PushFront("mca")
	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)
	fmt.Println(data.Front().Next().Value)
	fmt.Println(data.Back().Prev().Value)

	for i := data.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	for i := data.Back(); i != nil; i = i.Prev() {
		fmt.Println(i.Value)
	}
}

// go bulid hello-world.go
// go run hello-world.go
