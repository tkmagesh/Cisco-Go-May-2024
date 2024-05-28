/*
Refactoring to functions
*/
package main

import "fmt"

func main() {
	start, end := getRange()
	for no := start; no <= end; no++ {
		if isPrime(no) {
			fmt.Println("Prime No :", no)
		}
	}
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func getRange() (start, end int) {
	fmt.Println("Enter the range")
	fmt.Scanln(&start, &end)
	return
}
