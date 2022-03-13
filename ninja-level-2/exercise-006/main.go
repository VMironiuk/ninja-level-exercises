package main

import "fmt"

const (
	currentYear = 2022 + iota
	year2023    = currentYear + iota
	year2024    = currentYear + iota
	year2025    = currentYear + iota
	year2026    = currentYear + iota
)

func main() {
	fmt.Println(currentYear, year2023, year2024, year2025, year2026)
}
