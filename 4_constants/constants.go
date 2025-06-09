package main

import "fmt"

const age = 30

var str string = "hello"

// name = "raj" (this type cannot be declared outside the function)

func main() {
	// const name string = "golang"
	fmt.Println(age)

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}