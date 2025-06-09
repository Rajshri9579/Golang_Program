package main

import (
	"fmt"
	
)

//slices -> dynamic arrays
//most used construct in go
// + useful methods
func main() {
	//uninitialised slice is nil
	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	// var nums = make([]int, 0, 3)
	// //capacity ->> maximum number of elements that can be fit
	// //the capacity of slice gets double, when (capacity + 1) is reached
	// num := 1
	// for i := 0; i < 5; i++{
	// 	nums = append(nums, num)
	// 	num++
	// }
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))

	//copy function
	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))
	// copy(nums2, nums)
	// fmt.Println(nums, nums2)

	//slice opereator
	// var nums = []int{1, 2, 3}
	// fmt.Println(nums[1:])     //[start:end], (end is excluded)




	//slice package
	// var nums1 = []int{1, 2}
	// var nums2 = []int{1, 2}

	// fmt.Println(slices.Equal(nums1, nums2))   //returns bolean value



	var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums)
}
