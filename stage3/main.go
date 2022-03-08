package main

import (
	"fmt"
	"math"
)

// LogWithBase is a Custom Logarithm function -- it takes a value 'x' and the logarithmic base to use 'b'
func LogWithBase(x, b float64) float64 {
	return math.Log(x) / math.Log(b)
}

func main() {
	fmt.Println("What do you want to calculate?")
	fmt.Println("Type \"n\" for number of monthly payments,")
	fmt.Println("Type \"a\" for annuity monthly payment amount,")
	fmt.Println("Type \"p\" for loan principal:")

	var paymentType string
	fmt.Scanln(&paymentType)

	switch paymentType {

	case "n":
		var loanPrincipal float64
		fmt.Println("Enter the loan principal:")
		fmt.Scanln(&loanPrincipal)

		var monthlyPayment float64
		fmt.Println("Enter the monthly payment:")
		fmt.Scanln(&monthlyPayment)

		var interestRate float64
		fmt.Println("Enter the loan interest:")
		fmt.Scanln(&interestRate)

		nomInterestRate := interestRate / (12 * 100)
		monthsToPay := math.Ceil(LogWithBase(monthlyPayment/
			(monthlyPayment-nomInterestRate*loanPrincipal),
			1+nomInterestRate))

		years := math.Floor(monthsToPay / 12)
		months := math.Mod(monthsToPay, 12)

		if monthsToPay == 1 {
			fmt.Printf("It will take %d month to repay this loan!", int(monthsToPay))
		} else {
			fmt.Printf("It will take %d years and %d months to repay this loan!", int(years), int(months))
		}

	case "p":
		var annuityPayment float64
		fmt.Println("Enter the annuity payment:")
		fmt.Scanln(&annuityPayment)

		var periods int
		fmt.Println("Enter the number of periods:")
		fmt.Scanln(&periods)

		var interestRate float64
		fmt.Println("Enter the loan interest:")
		fmt.Scanln(&interestRate)

		nomInterestRate := interestRate / (12 * 100)
		loanPrincipal := annuityPayment / ((nomInterestRate * math.Pow(1+nomInterestRate, float64(periods))) /
			(math.Pow(1+nomInterestRate, float64(periods)) - 1))

		fmt.Printf("Your loan principal = %d!", int(math.Floor(loanPrincipal)))

	case "a":
		var loanPrincipal float64
		fmt.Println("Enter the loan principal:")
		fmt.Scanln(&loanPrincipal)

		var periods int
		fmt.Println("Enter the number of periods:")
		fmt.Scanln(&periods)

		var interestRate float64
		fmt.Println("Enter the loan interest:")
		fmt.Scanln(&interestRate)

		nomInterestRate := interestRate / (12 * 100)
		monthlyPayment := math.Ceil(loanPrincipal * nomInterestRate /
			(1 - math.Pow(1+nomInterestRate, float64(-periods))))

		fmt.Printf("Your monthly payment = %d!", int(monthlyPayment))
	}
}
