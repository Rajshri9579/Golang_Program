package main

import "fmt"

func sum(nums ...int) int {
	total :=  0

	for _, num := range nums{
		total = total + num
	}

	return total
}

func main() {

	nums := []int{1, 7, 8, 5}
	result := sum(nums...) // OR fmt.Println(1, 2, 3, 4)
	fmt.Println(result)
	// fmt.Println(1, 2, 3, 4)
}