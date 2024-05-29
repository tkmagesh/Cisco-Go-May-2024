package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divisor cannot be zero")

func main() {
	for {
		var divisor int
		fmt.Println("Enter the divisor :")
		fmt.Scanln(&divisor)
		q, r, err := divideClient(100, divisor)
		if err != nil {
			fmt.Println("error :", err)
			continue
		}
		fmt.Printf("dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
		break
	}
}

// wrappter to conver the panic into an error
func divideClient(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

// 3rd party library that results in a panic
func divide(x, y int) (quotient, remainder int) {
	if y == 0 {
		panic(ErrDivideByZero)
	}
	quotient, remainder = x/y, x%y
	return
}
