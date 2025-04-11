package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	/*
		Gaols
		1) validate user input
		=> show error message and exit if invalid input is provided
		no negative numbers
		not 0
		2) store calculated results into file

	*/
	for {
		var revenue, expenses, tax_rate, err = readValues()
		if err != nil {
			panic(err)
			return
		}
		var ebt = calculateEBT(revenue, expenses)
		var profit = calculateProfit(ebt, tax_rate)
		var ratio = calculateRatio(ebt, profit)

		print(ebt, profit, ratio)
		writeDataToFile(ebt, profit, ratio)
	}

}

func readValues() (float64, float64, float64, error) {
	var revenue, expenses, tax_rate float64
	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter tax: ")
	fmt.Scan(&tax_rate)
	if revenue <= 0 || expenses <= 0 || tax_rate <= 0 {
		return 0, 0, 0, errors.New("No negative numbers or 0")
	} else {
		return revenue, expenses, tax_rate, nil
	}

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

func writeDataToFile(ebt float64, profit float64, ratio float64) {
	results := fmt.Sprintf("EBT: %v\nProfit: %v\nRatio: %v", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}
