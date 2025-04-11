package main

import "fmt"

func main() {
	var revenue, expenses, tax_rate = readValues()
	var ebt = calculateEBT(revenue, expenses)
	var profit = calculateProfit(ebt, tax_rate)
	var ratio = calculateRatio(ebt, profit)

	print(ebt, profit, ratio)

}

func readValues() (float64, float64, float64) {
	var revenue, expenses, tax_rate float64
	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter tax: ")
	fmt.Scan(&tax_rate)
	return revenue, expenses, tax_rate
}

func calculateEBT(revenue float64, expenses float64) float64 {
	ebt := revenue - expenses
	return ebt
}

func calculateProfit(ebt float64, tax_rate float64) float64 {
	profit := ebt * (1 - tax_rate/100)
	return profit
}

func calculateRatio(ebt float64, profit float64) float64 {
	ratio := ebt / profit
	return ratio
}

func print(ebt float64, profit float64, ratio float64) {
	fmt.Printf("ebt: %v\nprofit: %v\nratio: %v\n ", ebt, profit, ratio)
}
