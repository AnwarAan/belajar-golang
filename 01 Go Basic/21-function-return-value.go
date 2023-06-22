package main

import "fmt"

func intro(name string) string {
	return "Hello " + name
}

func finalScore(score int) string {
	if score < 4 {
		return "bad"
	} else if score > 4 && score < 7 {
		return "no bad"
	} else {
		return "good"
	}
}

func main() {
	whoAreYou := intro("MCA")
	fmt.Println(whoAreYou)
	fmt.Println(finalScore(5))
}

// go bulid hello-world.go
// go run hello-world.go
