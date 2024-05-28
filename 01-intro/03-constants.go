package main

import "fmt"

func main() {
	/*
		const pi float64 = 3.14
		fmt.Println(pi)
	*/
	// type inference
	const pi = 3.14
	fmt.Println(pi)

	// constants can remain unused
	const unUsed = -9999
}
