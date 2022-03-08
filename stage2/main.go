package main

import (
	"fmt"
	"math"
)

func main() {
	var loanPrincipal float64
	fmt.Println("Enter the loan principal:")
	fmt.Scanln(&loanPrincipal)

	fmt.Println("What do you want to calculate?")
	fmt.Println("Type \"m\" - for number of monthly payments,")
	fmt.Println("Type \"p\" - for the monthly payment:")
	var paymentType string
	fmt.Scanln(&paymentType)

	switch paymentType {

	case "m":
		var monthlyPayment float64
		fmt.Println("Enter the monthly payment:")
		fmt.Scanln(&monthlyPayment)

		monthsToPay := int(math.Ceil(loanPrincipal / monthlyPayment))
		if monthsToPay == 1 {
			fmt.Printf("It will take %d month to repay the loan\n", monthsToPay)
		} else {
			fmt.Printf("It will take %d months to repay the loan\n", monthsToPay)
		}

	case "p":
		var months int
		fmt.Println("Enter the number of months:")
		fmt.Scanln(&months)

		monthlyPayment := loanPrincipal / float64(months)

		if math.Mod(monthlyPayment, 2) != 0 {
			monthlyPayment = math.Ceil(monthlyPayment)
			lastPayment := loanPrincipal - monthlyPayment*float64(months-1)
			fmt.Printf("Your monthly payment = %.2f and the last payment = %.2f\n", monthlyPayment, lastPayment)
		} else {
			fmt.Printf("Your monthly payment is %.2f.\n", monthlyPayment)
		}
	}
}
