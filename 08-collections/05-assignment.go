/*
Refactor the following using functions

isPrime()
genPrimes(start, end)
main() => accept the range and print the prime numbers
*/

package main

import "fmt"

func main() {
	var start, end int
	fmt.Println("Enter the start & end :")
	fmt.Scanln(&start, &end)
	primes := genPrimes(start, end)
	for _, no := range primes {
		fmt.Println("Prime No :", no)
	}
}

func genPrimes(start, end int) []int {
	var primes []int
	for i := start; i <= end; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
