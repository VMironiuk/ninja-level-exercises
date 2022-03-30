package main

import "fmt"

type person struct {
	fName string
	lName string
	age   int
}

func (p person) pSpeak() {
	fmt.Println("Hi, I am", p.fName, p.lName, "and I am", p.age, "years old.")
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (sa secretAgent) saSpeak() {
	sa.person.pSpeak()
	fmt.Println("... and I am a secret agent =]")
}

func main() {
	p := person{
		fName: "Miss",
		lName: "Moneypenny",
		age:   27,
	}
	fmt.Println(p.fName)
	p.pSpeak()

	sa := secretAgent{
		person: person{
			fName: "James",
			lName: "Bond",
			age:   32,
		},
		licenseToKill: true,
	}
	fmt.Println(sa.person.fName)
	sa.saSpeak()
	sa.person.pSpeak()
}
