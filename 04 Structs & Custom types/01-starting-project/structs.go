package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	dateAdd   time.Time
}

func main() {
	// firstName := getUserData("Please enter your first name: ")
	// lastName := getUserData("Please enter your last name: ")
	// birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user
	appUser.firstName = "Hussam"
	appUser.lastName = "Alghamdi"
	appUser.birthdate = "03/02/2002"

	fmt.Println(appUser)

	// ... do something awesome with that gathered data!

	// fmt.Println(firstName, lastName, birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
