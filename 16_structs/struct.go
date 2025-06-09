package main

import (
	"fmt"
	// "time"
)

//order struct

// type order struct {
// 	id string
// 	amount float32
// 	status string
// 	createdAt time.Time //nanosecond precision
// }

// func newOrder(id string, amount float32, status string) *order {
// 	//initial setup goes here...
// 	myOrder := order{
// 		id: id,
// 		amount: amount,
// 		status: status,
// 	}

// 	return &myOrder
// }

// //receiver type
// func (o *order) changeStatus(status string){
// 	o.status = status
// }

// func (o order) getAmount() float32 {
// 	return o.amount
// }

func main(){
	//if you don't set any field, default values is set
	// myOrder := order{
	// 	id: "1",
	// 	amount: 50.450,
	// 	status: "received",
	// } // we can create multiple instances of a struct

	
	// myOrder := newOrder("3", 55.00, "received")
	// fmt.Println(myOrder)
	// myOrder.createdAt = time.Now()
	// fmt.Println(myOrder.amount)
	// fmt.Println(myOrder.status)











	language := struct {
		name string
		isGood bool
	} {"golang", true}

	fmt.Println(language)
}