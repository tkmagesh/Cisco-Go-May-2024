package main

import (
	"errors"
	"fmt"
)

func main() {
	divisor := 7
	if q, r, err := divide(100, divisor); err == nil {
		fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
	} else {
		fmt.Println("error :", err)
	}
}

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		return 0, 0, errors.New("divisor cannot be zero")
	}
	quotient, remainder := x/y, x%y
	return quotient, remainder, nil
}
*/

// using named result
func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = errors.New("divisor cannot be zero")
		return
	}
	quotient, remainder = x/y, x%y
	return
}
