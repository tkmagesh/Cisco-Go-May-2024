package main

import "fmt"

func main() {
	/*
		1. User input cannot have any spaces
		2. It is assumed that the user will enter only valid data (error is returned in case of invalid data)
	*/
	/*
		var userName string
		fmt.Println("Enter your name :")
		fmt.Scanln(&userName)
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	*/

	var x, y int
	fmt.Scanln(&x, &y)
	fmt.Println(x + y)
}
