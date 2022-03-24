// Package to calculate dogs age.
package dog

// Count of dog years in one human year.
const dogYears = 7

// Years takes human years value and converts it to dog years value.
func Years(y int) int {
	return y * dogYears
}
