package main

import "fmt"

func main() {
	// var nos []int
	// var nos []int = []int{20, 40, 10, 30}
	// var nos = []int{20, 40, 10, 30}
	nos := []int{20, 40, 10, 30}
	fmt.Println(nos)

	nos = append(nos, 3)
	nos = append(nos, 1, 2, 5, 4)
	hundreds := []int{300, 100, 500}
	nos = append(nos, hundreds...)
	fmt.Printf("len(nos) = %d, nos = %v\n", len(nos), nos)

	fmt.Println("Iteration using index")
	for i := 0; i < len(nos); i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("Iteration using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	// the pointer to the underlying array is copied
	/*
		nos2 := nos
		nos2[0] = 9999
		fmt.Printf("nos[0] = %d, nos2[0] = %d\n", nos[0], nos2[0])
	*/

	// use copy() to create a copy
	fmt.Println("Using copy() to create a copy")
	newNos := make([]int, len(nos))
	copy(newNos, nos)
	newNos[0] = 9999
	fmt.Printf("nos[0] = %d, newNos[0] = %d\n", nos[0], newNos[0])

	fmt.Println("Slice sort")
	sort(nos)
	// fmt.Printf("%p\n", &nos)
	fmt.Println(nos)

	fmt.Println("slicing")
	fmt.Printf("nos[2:5] = %v\n", nos[2:5])
	fmt.Printf("nos[2:] = %v\n", nos[2:])
	fmt.Printf("nos[:5] = %v\n", nos[:5])

}

func sort(list []int) /* no return values */ {
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
