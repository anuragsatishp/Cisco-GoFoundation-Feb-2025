package main

import "fmt"

func main() {
	var productRanks map[string]int
	// productRanks = map[string]int{"pen": 5, "pencil": 1, "marker": 3}
	/*
		productRanks = map[string]int{
			"pen":    5,
			"pencil": 1,
			"marker": 3,
		}
	*/

	productRanks = make(map[string]int)
	productRanks["pen"] = 5
	productRanks["pencil"] = 1
	productRanks["marker"] = 3
	fmt.Println(productRanks)

	fmt.Println("len(productRanks) =", len(productRanks))

	fmt.Println("Iteration using range")
	for k, v := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", k, v)
	}

	fmt.Println("Check for the existance of a key")
	// keyToCheck := "scribble-pad"
	keyToCheck := "pen"
	if rank, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Rank of %q is %d\n", keyToCheck, rank)
	} else {
		fmt.Printf("key %q does not exist\n", keyToCheck)
	}

	fmt.Println("use delete() to remove an item")
	// keyToRemove := "pen"
	keyToRemove := "scribble-pad"
	delete(productRanks, keyToRemove)
	fmt.Printf("After removing %q, productRanks = %v\n", keyToRemove, productRanks)
}
