package main

import "fmt"

func main() {
	/*
		var userName string
		fmt.Println("Enter the username:")
		fmt.Scanln(&userName)
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	*/

	var x, y int
	fmt.Println("Enter the numbers (seperated by a space):")
	fmt.Scanln(&x, &y)
	fmt.Println(x + y)
}
