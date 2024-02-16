package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 4.3

	fmt.Print("Enter Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter Expected Return % Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter Amount of Investment years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Print("Future Value\n")
	fmt.Println(futureValue)

	fmt.Print("Future Real Value With Inflation Rate of 6.5\n")
	fmt.Println(futureRealValue)
}
