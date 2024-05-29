/*
Refactoring to genPrimes() which will return a collection of prime numbers
*/
package main

import "fmt"

func main() {
	start, end := getRange()
	primes := genPrimes(start, end)
	for _, primeNo := range primes {
		fmt.Println("Prime No :", primeNo)
	}
}

func genPrimes(start, end int) []int {
	var primes []int
	for no := start; no <= end; no++ {
		if isPrime(no) {
			primes = append(primes, no)
		}
	}
	return primes
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
