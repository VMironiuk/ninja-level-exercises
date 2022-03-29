package main

import "fmt"

func main() {
	age := 42
	if age < 30 {
		fmt.Println("Less than 30 years old")
	} else if age >= 30 && age <= 60 {
		fmt.Println("Less than or equal 60 but greater than 30 years old")
	} else {
		fmt.Println("Greater than 60 years")
	}
}
