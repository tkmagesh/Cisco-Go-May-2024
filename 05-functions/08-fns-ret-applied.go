package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// ver 1.0
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	/*
		logAdd := getLogOperation(add)
		logAdd(100, 200)

		logSubtract := getLogOperation(subtract)
		logSubtract(100, 200)

		getLogOperation(multiply)(100, 200)

		logDivide := getLogOperation(func(x, y int) {
			fmt.Println("Divide Result :", x/y)
		})
		logDivide(100, 20)
	*/

	logAdd := getLogOperation(add)
	profiledAdd := getProfiledOperation(logAdd)
	profiledAdd(100, 200)

}

func multiply(x, y int) {
	fmt.Println("Multiply result :", x*y)
}

// profile wrapper
func getProfiledOperation(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		operation(x, y)
		elapsed := time.Since(start)
		fmt.Println("Elapsed :", elapsed)
	}
}

// generic log wrapper
func getLogOperation(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		log.Println("Operation started")
		operation(x, y)
		log.Println("Operation completed")
	}
}

// 3rd party library (no privilege to change the code)
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
