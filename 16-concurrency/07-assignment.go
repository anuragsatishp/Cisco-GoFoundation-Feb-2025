/*
refactor this to exploit the go concurrency model
*/
package main

import "fmt"

func main() {
	var start, end int
	fmt.Println("Enter the start & end :")
	fmt.Scanln(&start, &end)
NO_LOOP:
	for no := start; no <= end; no++ {
		for i := 2; i <= (no / 2); i++ {
			if no%i == 0 {
				continue NO_LOOP
			}
		}
		fmt.Println("Prime No :", no)
	}
}
