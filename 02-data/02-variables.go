package main

import "fmt"

func main() {
	/*
		var userName string (userName is "" by default)
		userName = "Magesh"
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	*/

	/*
		var userName string = "Magesh"
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	*/

	// type inference (by the compiler)
	/*
		var userName = "Magesh"
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	*/

	// Using ":="
	userName := "Magesh"
	fmt.Printf("Hi %s, Have a nice day!\n", userName)

	var x int
}
