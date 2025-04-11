package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceAccount = "balance.txt"

func getBalanceFromFile() (float64, error) {
	money, err := os.ReadFile(accountBalanceAccount)
	if err != nil {
		return 1000, errors.New("Faield to find balance file.")
	}
	balanceTxt := string(money)
	myBank, err := strconv.ParseFloat(balanceTxt, 64)

	if err != nil {
		return 1000, errors.New("can not converte file data to digit")
	}
	return myBank, nil
}

func writeBalanceToFile(balance float64) {
	balanceTxt := fmt.Sprint(balance)
	os.WriteFile(accountBalanceAccount, []byte(balanceTxt), 0644)
}

func main() {
	var myBank, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("Error: ", err)
	}
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
