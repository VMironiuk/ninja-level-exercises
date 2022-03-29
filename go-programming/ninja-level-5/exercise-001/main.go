package main

import "fmt"

type person struct {
	firstName                string
	lastName                 string
	favouriteIceCreamFlavors []string
}

func main() {
	s := []person{}
	s = append(s, person{
		firstName: "James",
		lastName:  "Bond",
		favouriteIceCreamFlavors: []string{
			"Chockolate",
			"Martini",
		},
	})
	s = append(s, person{
		firstName: "Miss",
		lastName:  "Moneypenny",
		favouriteIceCreamFlavors: []string{
			"Vanilla",
			"Bananas",
		},
	})
	for _, v := range s {
		fmt.Printf("%v %v likes:\n", v.firstName, v.lastName)
		for _, f := range v.favouriteIceCreamFlavors {
			fmt.Printf("\t%v\n", f)
		}
	}
}
