package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// var nos [5]int // memory is allocated & initialized
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}
	// var nos [5]int = [5]int{3, 1, 4}

	// type inference
	// var nos = [5]int{3, 1, 4, 2, 5}
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	// len() returns the size of the array
	fmt.Println("len(nos) :", len(nos))

	// iterating using indexer
	fmt.Println("iterating using indexer")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Printf("nos[%d] = %d\n", idx, nos[idx])
	}

	// iterating using range
	fmt.Println("iterating using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	// Arrays are just values
	l1 := [5]int{3, 1, 4, 2, 5}
	l2 := [5]int{3, 1, 4, 2, 5}
	fmt.Println(l1 == l2)

	fmt.Println("Before sorting, nos = ", nos)
	sort(&nos)
	fmt.Println("After sorting, nos = ", nos)

	/*
		fmt.Printf("&nos : %p\n", &nos)
		fmt.Printf("&nos[1] : %p\n", &nos[1])
	*/

	// Pointer arithmatic using unsafe package
	ptr := unsafe.Pointer(&nos) // address of element 0
	p2 := unsafe.Add(ptr, 8)    // address of element 1
	fmt.Println(p2)
	fmt.Println(*(*int)(p2)) // value of element 1 using the pointer
}

func sort(list *[5]int) {
	fmt.Printf("list : %p\n", list)
	fmt.Printf("list[1] : %p\n", &list[1])
	/*
		fmt.Println((*list)[0])
		fmt.Println(list[0])
	*/

	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}

}
