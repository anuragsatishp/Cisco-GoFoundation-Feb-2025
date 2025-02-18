package main

import "fmt"

func main() {
	// var nos [5]int
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}
	// var nos = [5]int{3, 1, 4, 2, 5}
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("Iteration using Index")
	for i := 0; i < 5; i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("Iteration using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	// Assignment of an array will result in creating a copy (coz arrays are also values in Go)
	nos2 := nos
	nos2[0] = 9999
	fmt.Printf("nos[0] = %d, nos2 = %d\n", nos[0], nos2[0])

	// Arrays are compared by value
	newNos := [5]int{3, 1, 4, 2, 5}
	fmt.Println("nos == newNos ?:", nos == newNos)

	// Use pointers to create references
	var nosPtr *[5]int
	nosPtr = &nos
	// fmt.Println((*nosPtr)[0])
	fmt.Println(nosPtr[0])

	fmt.Println("Array sort")
	sort( /* ? */ )
	fmt.Println(nos)
}

func sort( /* ? */ ) /* no return values */ {
	// implement bubble sort
}
