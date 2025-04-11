package main

import "fmt"

func main() {
	age := 32
	fmt.Println("age", age)
	fmt.Println(getAdultYears(age))
	fmt.Println(age)
}

func getAdultYears(age int) int {
	return age - 18
}
