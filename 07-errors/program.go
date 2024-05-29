package main

import (
	"errors"
	"fmt"
	"math/rand"
)

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

	for range 20 {
		evenNo, err := getRandEvenNo()
		if err == nil {
			fmt.Println("Even number generated successfully :", evenNo)
			continue
		}
		fmt.Println("error :", err)
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
