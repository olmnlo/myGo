package main

import (
	"fmt"
	"os"
	"strconv"
)

func getBalanceFromFile() float64 {
	money, _ := os.ReadFile("balance.txt")
	balanceTxt := string(money)
	myBank, _ := strconv.ParseFloat(balanceTxt, 64)
	return myBank
}

func writeBalanceToFile(balance float64) {
	balanceTxt := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceTxt), 0644)
}

func main() {
	var myBank float64 = getBalanceFromFile()
	fmt.Println("Welcome to bank")
	for {
		fmt.Println("What do you want?")
		fmt.Printf("1. Check balance\n2. Deposit money\n3. Withdraw money\n4. Exit\n")
		var howMuch float64
		var choice int64
		fmt.Print("your chose num: ")
		fmt.Scan(&choice)
		if choice == 1 {
			fmt.Println("your account: ", myBank)
		} else if choice == 2 {
			fmt.Scan(&howMuch)
			if howMuch > myBank {
				fmt.Println("you can not do this")
			} else {
				myBank -= howMuch
				writeBalanceToFile(myBank)
				fmt.Println("your account: ", myBank)
			}
		} else if choice == 3 {
			fmt.Scan(&howMuch)
			myBank += howMuch
			writeBalanceToFile(myBank)
			fmt.Println("your account: ", myBank)
		} else {
			fmt.Println("Thank you for using my bank")
			break
		}
	}
}
