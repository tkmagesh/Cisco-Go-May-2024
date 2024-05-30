package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() // scheduling f1 through the scheduler (to be executed in future)
	f2()

	// block the main() so that the scheduler can look for other goroutines scheduled & execute them

	// DO NOT DO THIS IN PRODUCTION
	time.Sleep(1 * time.Second)

	// fmt.Scanln()
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
