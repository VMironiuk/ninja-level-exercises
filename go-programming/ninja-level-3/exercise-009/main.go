package main

import "fmt"

func main() {
	favSport := "Hockey"
	switch favSport {
	case "Football":
		fmt.Println("My favourite is Bayern Munich!")
	case "Tennis":
		fmt.Println("Go Djokovic!")
	case "Hockey":
		fmt.Println("Finland is the winner of Olympic Games 2022!")
	case "Basketball":
		fmt.Println("I play basketball like Jordan")
	}
}
