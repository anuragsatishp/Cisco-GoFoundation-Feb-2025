/*
refactor this to exploit the go concurrency model
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	var start, end int
	fmt.Println("Enter the start & end :")
	fmt.Scanln(&start, &end)
	ch := genPrimes(start, end)
	for primeNo := range ch {
		fmt.Println("Prime No :", primeNo)
	}
	fmt.Println("Done")
}

func genPrimes(start, end int) chan int {
	var ch chan int
	ch = make(chan int)
	go func() {
		wg := sync.WaitGroup{}
		for no := start; no <= end; no++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if isPrime(no) {
					ch <- no
				}
			}()
		}
		wg.Wait()
		close(ch)
	}()
	return ch
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
