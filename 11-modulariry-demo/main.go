package main

import (
	"fmt"

	calc "github.com/tkmagesh/cisco-go-may-2024/11-modularity-demo/calculator"
	"github.com/tkmagesh/cisco-go-may-2024/11-modularity-demo/calculator/utils" //nested package
)

var app_name = "Modularity-Demo"

func main() {
	run()
	/*
		fmt.Println(calculator.Add(100, 200))
		fmt.Println(calculator.Subtract(100, 200))
		fmt.Println("OpCount :", calculator.OpCount())
	*/

	fmt.Println(calc.Add(100, 200))
	fmt.Println(calc.Subtract(100, 200))
	fmt.Println("OpCount :", calc.OpCount())
	// using a nested package
	fmt.Println(utils.IsEven(21))
}
