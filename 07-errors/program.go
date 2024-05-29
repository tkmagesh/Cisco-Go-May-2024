package main

import (
	"errors"
	"fmt"
	"math/rand"
)

var ErrDivideByZero error = errors.New("divisor cannot be zero")

func main() {
	/*
		for range 20 {
			err := generateRandEvenNo()
			if err == nil {
				fmt.Println("Even number generated successfully")
				continue
			}
			fmt.Println("error :", err)
		}
	*/

	/*
		for range 20 {
			evenNo, err := getRandEvenNo()
			if err == nil {
				fmt.Println("Even number generated successfully :", evenNo)
				continue
			}
			fmt.Println("error :", err)
		}
	*/

	divisor := 0
	if q, r, err := divide(100, divisor); err == nil {
		fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
	} else if err == ErrDivideByZero {
		fmt.Println("do not attempt to divide by 0")
	} else {
		fmt.Println("unknown error :", err)
	}

}

func generateRandEvenNo() error {
	if no := rand.Intn(100); no%2 == 0 {
		return nil
	}
	// creating an error
	var err error
	err = errors.New("Unable to generate random even number")
	return err
}

func getRandEvenNo() (int, error) {
	if no := rand.Intn(100); no%2 == 0 {
		return no, nil
	}
	// creating an error
	var err error
	err = errors.New("Unable to generate random even number")
	return 0, err
}

func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = ErrDivideByZero
		return
	}
	quotient, remainder = x/y, x%y
	return
}
