package main

import "fmt"

func main() {
	// Create instances of different payment types
	cc := CreditCard{CardNumber: "1234-5678-9012-3456"}
	pp := PayPal{Email: "user@example.com"}

	// Process payments
	ProcessPayment(cc, 100.0)
	ProcessPayment(pp, 200.0)
}

type Payer interface {
	Pay(amount float64) string
}

// Type 1
type CreditCard struct {
	CardNumber string
}

func (cc CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using Credit Card %s", amount, cc.CardNumber)
}

// Type 2
type PayPal struct {
	Email string
}

func (pp PayPal) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using PayPal account %s", amount, pp.Email)
}

// A function that accepts ANY Payer
func ProcessPayment(p Payer, amount float64) {
	fmt.Println(p.Pay(amount)) // Evaluated at runtime
}
