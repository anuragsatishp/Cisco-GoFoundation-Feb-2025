package main

import "fmt"

var count int

func main() {
	for range 200 {
		increment()
	}
	fmt.Println("Count :", count)
}

func increment() {
	count++
}
