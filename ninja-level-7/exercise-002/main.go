package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	p := person{
		firstName: "James",
		lastName:  "Bond",
		age:       32,
	}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}

func changeMe(p *person) {
	p.firstName = "Miss"
	p.lastName = "Moneypenny"
	p.age = 27
}
