package main

import (
	"fmt"
	"math/rand"
)

type Product struct {
}

func main() {
	// var x interface{}
	var x any
	x = 100
	x = "Sunt esse sit nostrud reprehenderit cupidatat sit."
	x = true
	x = 99.99
	fmt.Println(x)

	x = 100
	// x = "Laboris in proident aute eiusmod aliquip labore esse laborum cupidatat eiusmod duis id."
	// x = getValue()

	// will not work
	// y := x * 2

	// works but not safe
	// y := x.(int) * 2

	// safe (type assertion)
	if val, ok := x.(int); ok {
		y := val * 2
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	fmt.Println("type assertion using 'type-switch'")
	// type assertion using "type-switch"
	// x = 100
	// x = "Excepteur labore consequat tempor ut anim."
	// x = true
	// x = 99.76
	// x = struct{}{}
	x = Product{}
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x * 2 =", val*2)
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case float64:
		fmt.Println("x is a float64, x * 0.99 =", val*0.99)
	case struct{}:
		fmt.Println("x is a zero byte struct")
	default:
		fmt.Println("x is of unknown type")
	}

}

func getValue() interface{} {
	if rand.Intn(200)%2 == 0 {
		return 100
	}
	return "Do magna do proident exercitation id amet tempor laborum incididunt pariatur proident nostrud ullamco."
}
