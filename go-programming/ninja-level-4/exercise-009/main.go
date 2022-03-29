package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice cream", "Sunsets"},
	}

	m["john_smith"] = []string{"Popular", "Chocolate"}

	for k, s := range m {
		fmt.Println("Record", k)
		for i, v := range s {
			fmt.Printf("\tindex: %v\tvalue: %v\n", i, v)
		}
	}
}
