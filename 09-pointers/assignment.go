package main

import "fmt"

func main() {
	n1, n2 := 100, 200
	fmt.Printf("Before swapping, n1 = %d and n2 = %d\n", n1, n2)
	swap(&n1, &n2)
	fmt.Printf("After swapping, n1 = %d and n2 = %d\n", n1, n2)

}

func swap(x, y *int) /* no return results */ {
	*x, *y = *y, *x
}
