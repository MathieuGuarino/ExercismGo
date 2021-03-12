// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	var divisibleBy4 bool = year%4 == 0
	var divisibleBy100 bool = year%100 == 0
	var divisibleBy400 bool = year%400 == 0
	return (divisibleBy4 && divisibleBy100 && divisibleBy400) || (divisibleBy4 && !divisibleBy100)
}
