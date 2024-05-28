/*
anonymouys functions
	- can't have name
	- should be created & executed immediately
*/

package main

import "fmt"

func main() {

	func() {
		fmt.Println("Hi!")
	}()

	func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}("Magesh")

	func(firstName, lastName string) {
		fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
	}("Magesh", "Kuppan")

	var message = func(firstName, lastName string) string {
		return fmt.Sprintf("Hi %s %s, Have a nice day!", firstName, lastName)
	}("Magesh", "Kuppan")
	fmt.Println("Message :", message)

	q, r := func(x, y int) (quotient, remainder int) {
		quotient = x / y
		remainder = x % y
		return
	}(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)

}
