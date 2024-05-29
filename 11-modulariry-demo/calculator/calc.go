package calculator

import "fmt"

// private
var opCount map[string]int

// package init function
func init() {
	fmt.Println("calculator[calc.go] initialized")
	opCount = make(map[string]int)
}

// public
func OpCount() map[string]int {
	return opCount
}
