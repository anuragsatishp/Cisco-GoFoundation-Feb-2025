package main

import "fmt"

func main() {
	sayHi()
	sayHello("Magesh")
	greet("Magesh", "Kuppan")
	fmt.Println(getGreetMsg("Suresh", "Kannan"))
}

func sayHi() {
	fmt.Println("Hi there!")
}

func sayHello(userName string) {
	fmt.Printf("Hello %s!\n", userName)
}

/*
func greet(firstName string, lastName string) {
	fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}
*/

func greet(firstName, lastName string) {
	fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}

func getGreetMsg(firstName, lastName string) string {
	return fmt.Sprintf("Hi %s %s, Have a good day!", firstName, lastName)
}
