package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	firstName string
	lastName  string
}

func (p *person) speak() {
	fmt.Printf("Hi, I'm %v %v\n", p.firstName, p.lastName)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		firstName: "James",
		lastName:  "Bond",
	}
	saySomething(&p)
}
