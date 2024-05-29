package calculator

import "fmt"

func init() {
	fmt.Println("calculator[add.go] initialized")
}

func Add(x, y int) int {
	opCount["add"]++
	return x + y
}
