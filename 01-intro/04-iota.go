package main

import "fmt"

func main() {

	/*
		const Red int = 0
		const Green int = 1
		const Blue int = 2
		fmt.Printf("Red = %d, Green = %d and Blue = %d\n", Red, Green, Blue)
	*/

	/*
		const Red = 0
		const Green = 1
		const Blue = 2
		fmt.Printf("Red = %d, Green = %d and Blue = %d\n", Red, Green, Blue)
	*/

	/*
		const (
			Red   int = 0
			Green int = 1
			Blue  int = 2
		)
		fmt.Printf("Red = %d, Green = %d and Blue = %d\n", Red, Green, Blue)
	*/

	/*
		const (
			Red int = iota
			Green
			Blue
		)
		fmt.Printf("Red = %d, Green = %d and Blue = %d\n", Red, Green, Blue)
	*/

	/*
		const (
			Red = iota
			Green
			Blue
		)
		fmt.Printf("Red = %d, Green = %d and Blue = %d\n", Red, Green, Blue)
	*/

	/*
		const (
			Red   = iota + 2 // 0 + 2
			Green            // 1 + 2
			Blue             // 2 + 2
		)
		fmt.Printf("Red = %d, Green = %d and Blue = %d\n", Red, Green, Blue)
	*/

	/*
		const (
			Red   = (iota * 2) + 1 // 0 + 1
			Green                  // 2 + 1
			Blue                   // 4 + 1
		)
		fmt.Printf("Red = %d, Green = %d and Blue = %d\n", Red, Green, Blue)
	*/

	const (
		Red   = (iota * 2) + 1 // 0 + 1
		Green                  // 2 + 1
		_
		Blue // 6 + 1
	)
	fmt.Printf("Red = %d, Green = %d and Blue = %d\n", Red, Green, Blue)

	// iota applied
	const (
		READ = iota << 1
		WRITE
		EXECUTE
	)
	fmt.Printf("READ = %d, WRITE = %d, EXECUTE = %d\n", READ, WRITE, EXECUTE)
	fmt.Printf("READ = %b, WRITE = %b, EXECUTE = %b\n", READ, WRITE, EXECUTE)

}
