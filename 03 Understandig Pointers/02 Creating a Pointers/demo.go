package main

import "fmt"

func main() {
	age := 32
	agePointer := &age
	fmt.Println("age", age)
	fmt.Println(getAdultYears(agePointer))
}

func getAdultYears(age *int) int {
	return *age - 18
}
