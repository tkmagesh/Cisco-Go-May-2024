package main

import "fmt"

func main() {

	// with 0 arguments & 0 results
	sayHi()

	// with 1 argument & 0 result
	greet("Magesh")

	// with 2 argument & 0 result
	greetUser("Magesh", "Kuppan")

	// with 2 argument & 1 result
	msg := getGreetMsg("Magesh", "Kuppan")
	fmt.Printf("Message : %s\n", msg)
}

func sayHi() {
	fmt.Println("Hi!")
}

func greet(userName string) {
	fmt.Printf("Hi %s, Have a nice day!\n", userName)
}

/*
func greetUser(firstName string, lastName string) {
	fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}
*/

func greetUser(firstName, lastName string) {
	fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}

func getGreetMsg(firstName, lastName string) string {
	return fmt.Sprintf("Hi %s %s, Have a nice day!", firstName, lastName)
}
