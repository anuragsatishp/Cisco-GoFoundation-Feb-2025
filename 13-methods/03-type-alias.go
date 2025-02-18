package main

import "fmt"

// alias for "string" type
type MyStr string

func (s MyStr) Length() int {
	return len(s)
}

func main() {
	str := MyStr("Elit consequat qui do amet dolore minim veniam sit non sunt. Voluptate quis consectetur elit enim est dolor. In laborum dolore commodo laborum aute est ut quis fugiat. Cupidatat esse reprehenderit est sit. Aute veniam dolor mollit ut ex duis pariatur do. Exercitation laborum nisi occaecat irure. Amet veniam cillum est sit laborum officia sint id reprehenderit adipisicing occaecat non non consequat.")
	fmt.Println(str.Length())
}
