package main

import "fmt"

func main() {
	exec(f1)
	exec(f2)
	exec(func() {
		fmt.Println("anon invoked")
	})

}

func exec(fn func()) {
	fn()
}

/*
func main() {
	exec("f1") //=> exec invokes f1
	exec("f2") //=> exec invokes f2

	exec("F1")
	exec("nonExistentFunc")
}

func exec(fnName string) {
	// execute a function based on the argument
	switch fnName {
	case "f1":
		f1()
	case "f2":
		f2()
	default:
		fmt.Println("invalid argument")
	}
}
*/

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
