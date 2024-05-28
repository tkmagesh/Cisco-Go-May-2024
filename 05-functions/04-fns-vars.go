/*
anonymouys functions
	- can't have name
	- should be created & executed immediately
*/

package main

import "fmt"

func main() {

	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi!")
	}
	sayHi()

	var greet func(string)
	greet = func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}
	greet("Magesh")

	var greetUser func(string, string)
	greetUser = func(firstName, lastName string) {
		fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
	}
	greetUser("Magesh", "Kuppan")

	var getGreetMsg func(string, string) string
	getGreetMsg = func(firstName, lastName string) string {
		return fmt.Sprintf("Hi %s %s, Have a nice day!", firstName, lastName)
	}
	var message = getGreetMsg("Magesh", "Kuppan")
	fmt.Println("Message :", message)

	var divide func(int, int) (int, int)
	divide = func(x, y int) (quotient, remainder int) {
		quotient = x / y
		remainder = x % y
		return
	}
	var q, r = divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)

}
