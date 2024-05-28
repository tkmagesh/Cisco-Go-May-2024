package main

import "fmt"

func main() {
	// fmt.Println(divide(100, 7))

	q, r := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)

	quotient, _ := divide(200, 3)
	fmt.Printf("Dividing 200 by 3, quotient = %d\n", quotient)
}

/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

// named results
func divide(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
