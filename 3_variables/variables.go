package main

import "fmt"

func main() {

	//first pproch
	var lang string = "go"
	fmt.Println(lang)

	var number int = 45
	fmt.Println(number)

	//another pproch (type is inferred automatically)
	var name = "golang"
	fmt.Println(name)

	//shorthand approach 
	age := 20
	fmt.Println(age)


	//one more 
	var girl string
	girl = "rajshri"
	fmt.Println(girl)

	var pi float32
	pi = 3.142
	fmt.Println(pi)
}