package main

import "fmt"

func main() {
	fmt.Println(divide(100, 7))
	/*
		q, r := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)
	*/
	// Use only the "quotient"
	/*
		q, _ := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d \n", q)
	*/
}

/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

// Named results (preferred if more than one result returned)
func divide(x, y int) (quotient, remainder int) {
	quotient, remainder = x/y, x%y
	return
}
