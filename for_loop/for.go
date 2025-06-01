package main

import "fmt"

//for -> only construct in go for looping
func main(){
	//while loop
	// i := 1
	// for i <= 3{
	// 	fmt.Println(i)
	// 	i = i + 1
	// }


	// for{
	// 	fmt.Println(69)
	// }


	//classic for loop
	for i := 0; i <= 3; i++ {
		if i == 2{
			continue
		}
		if i == 3{
			break
		}
		fmt.Println(i)
	}

	for i := range 3 {
		fmt.Println(i)//will start from 0 & 3 will be excluded
	}
}