// this is go file
package main

import "fmt" //formatter module



func main() {
	//Using var
	var name string = "Alex"
	//datatype declaration done below
	var age int32 = 37
	//as this is var, we can redefine, unlike const
	var isCool = true

	//shorthand definition
	name2 := "Lewis"
	//shorthand for float64
	size := 1.3

	isCool = false

	//declaring multiple variables
	name, email := "Alex", "alex@gmail.com"

	fmt.Println(name, age, isCool, name2, email, size)
	fmt.Printf("%T\n",email)
}
