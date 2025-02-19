package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := &sync.WaitGroup{}
	wg.Add(10)
	go f1(wg)
	wg.Wait() // block the execution of this function until the wg counter becomes 0 (default = 0)

}

func f1(wg *sync.WaitGroup) {
	defer wg.Done() //decrement the counter by 1
	fmt.Println("f1 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 completed")
}
