package main

import "fmt"

type person struct {
	firstName        string
	lastName         string
	favouriteFlavors []string
}

func main() {
	bond := person{
		firstName: "James",
		lastName:  "Bond",
		favouriteFlavors: []string{
			"Chockolate",
			"Martini",
		},
	}
	moneypenny := person{
		firstName: "Miss",
		lastName:  "Moneypenny",
		favouriteFlavors: []string{
			"Vanilla",
			"Bananas",
		},
	}
	m := map[string]person{
		bond.lastName:       bond,
		moneypenny.lastName: moneypenny,
	}
	for _, v := range m {
		fmt.Printf("%v %v likes:\n", v.firstName, v.lastName)
		for _, f := range v.favouriteFlavors {
			fmt.Printf("\t%v\n", f)
		}
	}
}
