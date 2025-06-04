package main

import (
	"fmt"
	"maps"
)

//maps(go) --> hash, object(js), dictionary(python)
func main(){
	//creating map
	// m := make(map[string]string)

	//setting an element
	// m["name"] = "golang"
	// m["area"] = "backend"


	//getting an element
	// fmt.Println(m["name"], m["area"])
	//IMP: if key values does not exist in map then it returns zero value

	// m := make(map[string]int)
	// m["age"] = 30
	// m["price"] = 400

	// // fmt.Println(m["age1"])
	// fmt.Println(len(m))

	// // delete(m, "age")
	// clear(m)
	// fmt.Println(m)



	// m := map[string]int{"price": 30, "phones": 3,}
	// fmt.Println(m)


	// m := map[string]int{"price": 30, "phones": 3,}
	// k, ok := m["price"]
	// fmt.Println(k)
	// if ok{
	// 	fmt.Println("all ok")
	// } else{
	// 	fmt.Println("not ok")
	// }


	m1 := map[string]int{"price": 50, "phones": 6}
	m2 := map[string]int{"price": 50, "phones": 6}
	fmt.Println(maps.Equal(m1, m2))
}