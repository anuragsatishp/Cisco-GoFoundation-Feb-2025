package main

import (
	"errors"
	"fmt"
)

var ErrZeroDivisor error = errors.New("divisor cannot be 0")
var ErrZeroMultiplier error = errors.New("multiplier cannot be 0")

func main() {
	// x := 0
	x := 100

	divisor := 7
	// divisor := 0
	if q, r, err := divide(x, divisor); err == nil {
		fmt.Println(q, r)
	} else {
		switch err {
		case ErrZeroDivisor:
			fmt.Println("Do not attempt to divide by 0")
		case ErrZeroMultiplier:
			fmt.Println("Make sure multiplier is a non zero value")
		default:
			fmt.Println("Unknown error :", err)
		}
	}

}

/*
func divide(x, divisor int) (int, int, error) {
	if divisor != 0 {
		quotient, remainder := x/divisor, x%divisor
		return quotient, remainder, nil
	}
	return 0, 0, errors.New("divisor cannot be 0")
}
*/

/*
func divide(x, divisor int) (quotient, remainder int, err error) {
	if divisor != 0 && x != 0 {
		quotient, remainder = x/divisor, x%divisor
		return
	}
	if divisor == 0 {
		err = errors.New("divisor cannot be 0")
		return
	}
	// x IS 0
	err = errors.New("multiplier cannot be 0")
	return
}
*/

func divide(x, divisor int) (quotient, remainder int, err error) {
	if divisor != 0 && x != 0 {
		quotient, remainder = x/divisor, x%divisor
		return
	}
	if divisor == 0 {
		err = ErrZeroDivisor
		return
	}
	// x IS 0
	err = ErrZeroMultiplier
	return
}
