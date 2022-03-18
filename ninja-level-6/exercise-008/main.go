package main

import "fmt"

func main() {
	f := funcFactory()
	f()
}

func funcFactory() func() {
	return func() {
		fmt.Println("Made from funcFactory")
	}
}
