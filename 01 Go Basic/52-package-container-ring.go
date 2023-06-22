package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	var data *ring.Ring = ring.New(5)
	data.value = "mca"

	for i := 0; i < data.Len(); i++ {
		data.Value = "value" + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	data.Do(func(value interface{}) {
		fmt.Println(value)
	})

	fmt.Println("")

	fmt.Println(*data)
}

// go bulid hello-world.go
// go run hello-world.go
