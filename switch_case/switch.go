package main

import (
	"fmt"
	
)

func main() {
	//simple switch
	// i := 3
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// default:
	// 	fmt.Println("three")
	// }




	//multiple condition switch
	// switch time.Now().Weekday(){
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("its weekend")
	// default:
	// 	fmt.Println("its workday")
	// }






	//type switch
	whoAmI := func (i interface{})  {
		switch t := i.(type){
		case int:
			fmt.Println("its an integer")
		case string:
			fmt.Println("its a string")
		case bool:
			fmt.Println("its a boolean")
		default:
			fmt.Println("other", t)
		}
	}


	whoAmI(3.56)
}