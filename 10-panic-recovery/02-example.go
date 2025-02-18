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
	var divisor int
	for {
		fmt.Println("Enter the divisor")
		fmt.Scanln(&divisor)
		if q, r, err := divideAdapter(100, divisor); err != nil {
			fmt.Println("Error :", err)
			fmt.Println("Try again!")
			continue
		} else {
			fmt.Println(q, r)
			break
		}
	}
}

func divideAdapter(x, y int) (q, r int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	q, r = divide(x, y)
	return
}

// OSS (code changes not allowed)
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
