package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (b ByAge) Len() int           { return len(b) }
func (b ByAge) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByAge) Less(i, j int) bool { return b[i].Age < b[j].Age }

type ByLast []user

func (b ByLast) Len() int           { return len(b) }
func (b ByLast) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByLast) Less(i, j int) bool { return b[i].Last < b[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}
	for _, u := range users {
		fmt.Println(u.First, u.Last, ":")
		sort.Strings(u.Sayings)
		for _, s := range u.Sayings {
			fmt.Printf("\t%v\n", s)
		}
	}

	sort.Sort(ByAge(users))
	fmt.Println("------------ Sorted by age:")
	for _, u := range users {
		fmt.Println(u.First, u.Last, ":")
		sort.Strings(u.Sayings)
		for _, s := range u.Sayings {
			fmt.Printf("\t%v\n", s)
		}
	}

	sort.Sort(ByLast(users))
	fmt.Println("------------ Sorted by last:")
	for _, u := range users {
		fmt.Println(u.First, u.Last, ":")
		sort.Strings(u.Sayings)
		for _, s := range u.Sayings {
			fmt.Printf("\t%v\n", s)
		}
	}

	fmt.Println("------------ Sorted sayings:")
	for _, u := range users {
		fmt.Println(u.First, u.Last, ":")
		sort.Strings(u.Sayings)
		for _, s := range u.Sayings {
			fmt.Printf("\t%v\n", s)
		}
	}
}
