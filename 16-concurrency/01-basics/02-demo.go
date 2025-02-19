package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() // scheduling the exec of this function through the scheduler
	f2()

	// DO NOT DO THIS (poor man's synchronization techniques)
	time.Sleep(4 * time.Second)

}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
