package main

import "fmt"

func main() {
	p := struct {
		locales []string
		dict    map[string]string
	}{
		locales: []string{"en", "de", "fr", "us"},
		dict:    map[string]string{"en": "England", "de": "Germany", "fr": "France", "us": "USA"},
	}
	fmt.Println(p)
	for i, v := range p.locales {
		fmt.Printf("%v %v\n", i, v)
	}
	for k, v := range p.dict {
		fmt.Printf("%v: %v\n", k, v)
	}
}
