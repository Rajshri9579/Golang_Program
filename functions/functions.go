package main

// import "fmt"

// func add(a int, b int) int { //or (a, b int)
// 	return a + b
// }

// func getLanguages() (string, string, bool) {
// 	return "golang", "javascript", false	
// }


// func processIt(fn func(a string ) string )  {
// 	fn("tanu")
// }


func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}

func main() {

	fn := processIt()
	fn(7)

	// fn := func (a string) string {
	// 	return "rajshri"
	// }
	// processIt(fn)

	// lang1, lang2, val := getLanguages()
	// fmt.Println(lang1, lang2, val)



	// result := add(3, 8)
	// fmt.Println(result)
}