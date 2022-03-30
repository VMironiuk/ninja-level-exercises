package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	fname string
	lname string
	age   int
}

func (p person) speak() {
	fmt.Println("Hi, I am", p.fname, p.lname, "and I am", p.age, "years old.")
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (sa secretAgent) speak() {
	sa.person.speak()
	if sa.licenseToKill {
		fmt.Println("I have a license to kill")
	} else {
		fmt.Println("I have not a license to kill")
	}
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		fname: "Miss",
		lname: "Moneypenny",
		age:   27,
	}

	sa := secretAgent{
		person: person{
			fname: "James",
			lname: "Bond",
			age:   32,
		},
		licenseToKill: true,
	}

	saySomething(p)
	saySomething(sa)
}
