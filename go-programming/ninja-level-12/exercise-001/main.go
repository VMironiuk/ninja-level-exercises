package main

import (
	"fmt"
	"vmyroniuk/dog"
)

type doggy struct {
	name string
	age  int
}

func main() {
	d := doggy{
		name: "Fido",
		age:  dog.Years(3),
	}
	fmt.Println(d)
}
