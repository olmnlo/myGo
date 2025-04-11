package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64 = 10
	const infaltionRate = 6.5

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow((1+infaltionRate/100), years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
