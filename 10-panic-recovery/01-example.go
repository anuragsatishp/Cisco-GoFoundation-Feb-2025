package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var ErrZeroDivisor error = errors.New("divisor cannot be 0")

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("App paniced!, err :", err)
			os.Exit(2)
		}
		fmt.Println("Thank you!")
	}()
	divisor := 0
	q, r := divide(100, divisor)
	fmt.Println(q, r)
}

func divide(x, y int) (q, r int) {
	log.Println("[divide] - calculating quotient")
	if y == 0 {
		panic(ErrZeroDivisor)
	}
	q = x / y
	log.Println("[divide] - calculating remainder")
	r = x % y
	return
}
