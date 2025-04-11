package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5
	var years float64 = 10
	const infaltionRate = 6.5

	futureRealValue, futureValue := calculate(investmentAmount, expectedReturnRate, years, infaltionRate)

	fmt.Printf("futureRealValue: %.3f\nfutureValue: %.3f", futureRealValue, futureValue)

}

func calculate(investmentAmount float64, expectedReturnRate float64, years float64, infaltionRate float64) (futureRealValue float64, futureValue float64) {
	futureValue = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue = futureValue / math.Pow((1+infaltionRate/100), years)
	return
}
