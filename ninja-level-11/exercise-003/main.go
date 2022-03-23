package main

import "fmt"

type customErr struct{}

func (ce customErr) Error() string {
	return "custom_error"
}

func main() {
	cerr := customErr{}
	foo(cerr)
}

func foo(err error) {
	fmt.Println(err.Error())
}
