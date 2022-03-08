package main

import (
	"flag"
	"fmt"
	"math"
)

// LogWithBase is a Custom Logarithm function -- it takes a value 'x' and the logarithmic base to use 'b'
func LogWithBase(x, b float64) float64 {
	return math.Log(x) / math.Log(b)
}

func main() {
	t := flag.String("type", "", "The type of payment -- enter 'annuity' or 'diff'")

	p := flag.Int("principal", 0, "Enter the principal amount")
	n := flag.Int("periods", 0, "Enter the number of periods")

	interest := flag.Float64("interest", 0, "Enter the interest rate")
	a := flag.Int("payment", 0, "Enter the payment amount")

	flag.Parse() // Remember to call this function to parse the flags!

	if t == nil || (*t == "diff" && *a != 0) || (*interest == 0) || (*n != 0 && *n < 0) {
		fmt.Println("Incorrect parameters")
	} else {
		// Create pointers to the previously created flags
		t_ := *t          // t_ is a pointer to the value of the flag 't'
		p_ := float64(*p) // p_ is a pointer to the value of the flag 'p'
		n_ := float64(*n) // n_ is a pointer to the value of the flag 'n'
		a_ := float64(*a) // a_ is a pointer to the value of the flag 'a'

		nomInterestRate := *interest / (12 * 100)

		if t_ == "diff" {
			overpayment := p_
			for i := 1; i <= *n; i++ {
				d1 := math.Ceil((p_ / n_) + nomInterestRate*(p_-(p_*(float64(i)-1))/n_))
				overpayment -= d1
				fmt.Printf("Month %d: payment is %d\n", i, int(d1))
			}
			fmt.Printf("Overpayment = %d\n", int(math.Abs(overpayment)))

		} else if t_ == "annuity" {
			if n_ == 0 {
				n_ = math.Ceil(LogWithBase(a_/(a_-nomInterestRate*p_), 1+nomInterestRate))
				years := math.Floor(n_ / 12)
				months := math.Mod(n_, 12)
				fmt.Printf("You need %d years and %d months to repay this credit!\n", int(years), int(months))
			} else if *p == 0 {
				p_ = math.Floor(a_ / (nomInterestRate / (1 - 1/math.Pow(1+nomInterestRate, n_))))
				fmt.Printf("Your credit principal = %d!\n", int(p_))
			} else {
				a_ = math.Ceil(p_ * nomInterestRate / (1 - 1/math.Pow(1+nomInterestRate, n_)))
				fmt.Printf("Your annuity payment = %d!\n", int(a_))
			}

			overpayment := a_*n_ - p_
			fmt.Printf("\nOverpayment = %d\n", int(math.Abs(overpayment)))
		}
	}
}
