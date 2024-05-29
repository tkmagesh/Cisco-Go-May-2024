package main

import "fmt"

func main() {
	var no int
	no = 100

	var noPtr *int
	noPtr = &no // address of no
	fmt.Println(noPtr)

	// accessing the underlying value using the pointer (deferencing)
	fmt.Println(*noPtr)
	fmt.Println(no == *(&no))

	// pointer applied
	increment(&no)
	fmt.Println("After incrementing")
	fmt.Println(no)
}

func increment(x *int) {
	fmt.Println("[increment] &x =", x)
	*x = *x + 1
}
