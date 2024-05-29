package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("app panicked : ", err)
			return
		}
		log.Println("app shutdown normal")
	}()
	for range 20 {
		evenNo := getRandEvenNo()
		fmt.Println("Even number generated successfully :", evenNo)

	}
}

func getRandEvenNo() int {
	if no := rand.Intn(100); no%2 == 0 {
		return no
	}
	// create a panic
	var err error
	err = errors.New("Unable to generate random even number")
	panic(err)
}
