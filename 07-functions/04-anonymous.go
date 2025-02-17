package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hi there!")
	}()

	func(userName string) {
		fmt.Printf("Hello %s!\n", userName)
	}("Magesh")

	func(firstName, lastName string) {
		fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
	}("Magesh", "Kuppan")

	s := func(firstName, lastName string) string {
		return fmt.Sprintf("Hi %s %s, Have a good day!", firstName, lastName)
	}("Suresh", "Kannan")
	fmt.Println(s)
}
