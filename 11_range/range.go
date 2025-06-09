package main

import "fmt"

//range --> iterating over data structures
func main(){
	// nums := []int{6, 7, 8}

	// for i := 0; i <len(nums); i++{
	// 	fmt.Println(nums[i])
	// }

	// sum := 0
	// for i, num := range nums{
	// 	 fmt.Println(num, i)
	// 	 fmt.Println(i)
	// 	 sum = sum + num	
	// }
	// fmt.Println(sum)


	// m := map[string]string{"fname": "john", "lname": "doe"}

	// for name, val := range m{
	// 	fmt.Println(name,":", val)
	// }

	// unicode code point rune(data structure)
	//first parameter(starting byte of rune)
	//255 -> 1 byte, 
	for i, j := range "rajshri"{
		fmt.Println(i, j, string(j))
	}
}