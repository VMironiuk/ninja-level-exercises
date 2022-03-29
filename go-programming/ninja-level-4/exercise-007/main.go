package main

import "fmt"

func main() {
	ss := [][]string{}
	ss = append(ss, []string{"James", "Bond", "Shaken, not stirred"})
	ss = append(ss, []string{"Miss", "Moneypenny", "Helloooooo, James."})

	for i, s := range ss {
		fmt.Println(i, s)
		for j, v := range s {
			fmt.Println(j, v)
		}
	}
}
