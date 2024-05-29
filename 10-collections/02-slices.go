package main

import "fmt"

func main() {
	// var nos []int
	// var nos []int = []int{3, 1, 4, 2, 5}
	// var nos []int = []int{3, 1, 4}

	// type inference
	// var nos = []int{3, 1, 4, 2, 5}
	nos := []int{3, 1, 4, 2, 5}
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

	// slices are pointer to underlying array
	/*
		l1 := []int{3, 1, 4, 2, 5}
		l2 := []int{3, 1, 4, 2, 5}
		fmt.Println(l1 == l2)  // slices can ONLY be compared to nil
	*/

	x1 := []int{3, 1, 4, 2, 5}
	x2 := x1    // create a copy of x1
	x2[0] = 100 // changing x2 Will affect x1
	fmt.Printf("x1 = %d, x2 = %d\n", x1, x2)

	// passing around a slice
	fmt.Println("Before sorting, nos = ", nos)
	sort(nos)
	fmt.Println("After sorting, nos = ", nos)

	// adding new values
	nos = append(nos, 10)
	fmt.Println("nos :", nos)

	nos = append(nos, 20, 30, 40)
	fmt.Println("nos :", nos)

	// append another slice
	hundreds := []int{100, 200, 300}
	nos = append(nos, hundreds...)
	fmt.Println("nos :", nos)

	// creating sub-slices
	fmt.Println("// creating sub-slices")
	fmt.Println("nos[2:6] = ", nos[2:6])
	fmt.Println("nos[2:] = ", nos[2:])
	fmt.Println("nos[:6] = ", nos[:6])
}

func sort(list []int) {
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
