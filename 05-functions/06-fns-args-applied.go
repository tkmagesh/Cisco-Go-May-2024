package main

import (
	"fmt"
	"log"
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

	logOperation(add, 100, 200)
	logOperation(subtract, 100, 200)
	logOperation(multiply, 100, 200)
	logOperation(func(x, y int) {
		fmt.Println("Divide Result :", x/y)
	}, 100, 20)
}

func multiply(x, y int) {
	fmt.Println("Multiply result :", x*y)
}

// generic log wrapper
func logOperation(operation func(int, int), x, y int) {
	log.Println("Operation started")
	operation(x, y)
	log.Println("Operation completed")
}

// log wrappers
func logAdd(x, y int) {
	log.Println("Operation started")
	add(x, y)
	log.Println("Operation completed")
}

func logSubtract(x, y int) {
	log.Println("Operation started")
	subtract(x, y)
	log.Println("Operation completed")
}

// 3rd party library (no privilege to change the code)
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
