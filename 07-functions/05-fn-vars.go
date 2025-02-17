package main

import "fmt"

func main() {
	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi there!")
	}
	sayHi()

	var sayHello func(string)
	sayHello = func(userName string) {
		fmt.Printf("Hello %s!\n", userName)
	}
	sayHello("Magesh")

	var greet func(string, string)
	greet = func(firstName, lastName string) {
		fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
	}
	greet("Magesh", "Kuppan")

	var getGreetMsg func(string, string) string
	getGreetMsg = func(firstName, lastName string) string {
		return fmt.Sprintf("Hi %s %s, Have a good day!", firstName, lastName)
	}
	s := getGreetMsg("Suresh", "Kannan")
	fmt.Println(s)

	var oper func(int, int)

	oper = func(i1, i2 int) {
		fmt.Println("Add :", i1+i2)
	}
	oper(100, 200)

	oper = func(i1, i2 int) {
		fmt.Println("Multiply :", i1*i2)
	}
	oper(100, 200)
}
