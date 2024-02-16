package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Print("Earnings before Tax\n")
	fmt.Println(ebt)
	fmt.Print("Your Profit after Tax\n")
	fmt.Println(profit)
	fmt.Print("Ratio\n")
	fmt.Println(ratio)
}
