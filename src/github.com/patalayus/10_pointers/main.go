// this is go file
package main

import "fmt"

func main() {
	//integer
	a := 5

	//memory location
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	//use * to read val from address
	fmt.Println(*b)

	//change val with pointer
	*b = 10
	fmt.Println(a)
}