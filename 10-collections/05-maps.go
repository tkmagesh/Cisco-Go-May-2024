package main

import "fmt"

func main() {
	// collection of products & and their ranks
	// var productRanks map[string]int
	// var productRanks map[string]int = map[string]int{"pen": 3, "pencil": 1, "marker": 5}

	// type inference
	// var productRanks = map[string]int{"pen": 3, "pencil": 1, "marker": 5}
	// productRanks := map[string]int{"pen": 3, "pencil": 1, "marker": 5}
	/*
		productRanks := map[string]int{
			"pen":    3,
			"pencil": 1,
			"marker": 5, // last item should also have ','
		}
	*/

	productRanks := make(map[string]int)
	productRanks["pen"] = 3
	productRanks["pencil"] = 1
	productRanks["marker"] = 5
	fmt.Println(productRanks)

	// iteration using range
	fmt.Println("// iteration using range")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	// check for the existence of a key
	fmt.Println("// check for the existence of a key")
	// keyToCheck := "scribble-pad"
	keyToCheck := "pencil"
	if rank, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Rank of %q = %d\n", keyToCheck, rank)
	} else {
		fmt.Printf("product %q does not exist\n", keyToCheck)
	}

	// removing a key/value pair
	// keyToRemove := "pen"
	keyToRemove := "scribble-pad"
	delete(productRanks, keyToRemove)
	fmt.Printf("After removing %q, productRanks = %v\n", keyToRemove, productRanks)
}
