package main

import "fmt"

func main() {
	var no int
	no = 100

	var noPtr *int
	noPtr = &no

	fmt.Println(noPtr, no)

	// deferencing (pointer -> value)
	fmt.Println(*noPtr)

	fmt.Println(no == *(&no))

	*noPtr = 200
	fmt.Println(no)
}
