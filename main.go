package main

import "fmt"

func main() {
	//simple request
	m := &payRequest{
		APIKey:       APIKEY,
		Amount:       50000,
		Redirect:     RedirectTo,
		FactorNumber: "50",
	}
	n := m.RequestPaymentToken()
	fmt.Println(n)

	//simple verify
	f := &verifyRequest{
		APIKEY:  APIKEY,
		TransID: 300,
	}
	l := f.RequestPaymentVerification()
	fmt.Println(l)
}
