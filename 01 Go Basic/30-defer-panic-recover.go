package main

import "fmt"

func logging() {
	fmt.Println("login")
}

func runApp(value int) {
	defer logging()
	fmt.Print("run app")
	result := 10 / value
	fmt.Println("result", result)
}

func endApp() {
	fmt.Println("end app")
	message := recover()
	if message != nil {
		fmt.Println("err", message)
	}
}

func boolApp(err bool) {
	defer endApp()
	if err {
		panic("err occure")
	}
	fmt.Println("run app")
}

func main() {
	//defer --> call func despite err
	//panic --> stop program
	//recover --> panic can't stop func

	// runApp(0)
	boolApp(true)

}

// go bulid hello-world.go
// go run hello-world.go
