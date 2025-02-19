package main

import (
	"fmt"
)

// share memory by communicating

func main() {
	ch := make(chan int)
	go func() {
		result := add(100, 200)
		ch <- result
	}()
	result := <-ch
	fmt.Println(result)
}

func add(x, y int) int {
	result := x + y
	return result
}
