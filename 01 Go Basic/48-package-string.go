package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("muchamad choirul anwar")
	fmt.Println(strings.Contains("muchamad choirul anwar", "anwar"))
	fmt.Println(strings.Split("muchamad choirul anwar", ""))
	fmt.Println(strings.ToLower("Muchamad Choirul Anwar"))
	fmt.Println(strings.ToUpper("muchamad choirul anwar"))
	fmt.Println(strings.ToTitle("muchamad choirul anwar"))
	fmt.Println(strings.Trim("     muchamad choirul anwar      ", " "))
	fmt.Println(strings.ReplaceAll("muchamad choirul anwar", "anwar", "mca"))
}

// go bulid hello-world.go
// go run hello-world.go
