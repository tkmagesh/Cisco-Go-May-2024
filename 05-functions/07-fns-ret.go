package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		time.Sleep(300 * time.Millisecond)
		fn := getFn()
		fn()
	}
}

func getFn() func() {
	if no := rand.Intn(100); no%2 == 0 {
		return f1
	}
	return f2
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
