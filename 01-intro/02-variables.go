package main

import "fmt"

func main() {
	/*
		var userName string
		userName = "Magesh"
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	*/

	/*
		var userName string = "Magesh"
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	*/

	// type inference
	/*
		var userName = "Magesh"
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	*/

	// using ":="
	userName := "Magesh"
	fmt.Printf("Hi %s, Have a nice day!\n", userName)

	/*
		var x int
		var y int
		var msg string
		var result int
		x = 100
		y = 200
		msg = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(msg, x, y, result)
	*/

	/*
		var x, y, result int
		var msg string
		x = 100
		y = 200
		msg = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(msg, x, y, result)
	*/

	/*
		var (
			x, y, result int
			msg          string
		)
		x = 100
		y = 200
		msg = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(msg, x, y, result)
	*/

	/*
		var x int = 100
		var y int = 200
		var msg string = "Sum of %d and %d is %d\n"
		var result int = x + y
		fmt.Printf(msg, x, y, result)
	*/

	/*
		var x, y int = 100, 200
		var msg string = "Sum of %d and %d is %d\n"
		var result int = x + y
		fmt.Printf(msg, x, y, result)
	*/

	/*
		var (
			x, y   int    = 100, 200
			msg    string = "Sum of %d and %d is %d\n"
			result int    = x + y
		)
		fmt.Printf(msg, x, y, result)
	*/

	// type inference
	/*
		var (
			x, y   = 100, 200
			msg    = "Sum of %d and %d is %d\n"
			result = x + y
		)
		fmt.Printf(msg, x, y, result)
	*/

	/*
		var (
			x, y, msg = 100, 200, "Sum of %d and %d is %d\n"
			result    = x + y
		)
		fmt.Printf(msg, x, y, result)
	*/

	// Using ":="

	x, y, msg := 100, 200, "Sum of %d and %d is %d\n"
	result := x + y
	fmt.Printf(msg, x, y, result)

}
