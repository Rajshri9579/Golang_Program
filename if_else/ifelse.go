package main

import "fmt"

func main() {
	// age := 90
	// if age == 18 {
	// 	fmt.Println("person is now an adult")
	// } else if age >= 80{
	// 	fmt.Println("person is a senior citizen")
	// } else {
	// 	fmt.Println("person is not an adult")
	// }



	var role = "admin"
	var hasPermission = true
	if role == "admin" || hasPermission {
		fmt.Printf("user is allowed an access")
	}
}