package main

import "fmt"

type payment struct {
	gateway razorpay
	
}

func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amount)

	// stripePaymentGw := stripe{}
	// 	p.gateway.payUsingStripe(amount)
	p.gateway.payUsingRazorpay(amount)

}

type razorpay struct {}

type stripe struct{}

func (s stripe) payUsingStripe(amount float32){
	fmt.Println("making payment using stripe", amount)
}

func (r razorpay) payUsingRazorpay(amount float32) {
	//logic to make payment
	fmt.Println("making payment using razorpay", amount)
}

func main() {
	newPayment := payment{}
	newPayment.makePayment(50000)
}