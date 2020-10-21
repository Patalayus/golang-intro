// this is go file
package main

import "fmt"

func main() {
	x := 10
	y := 10

	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	//else if

	colour := "yellow"

	if colour == "red" {
		fmt.Println("colour is red")
	} else if colour == "blue" {
		fmt.Println("colour is blue")
	} else {
		fmt.Println("colour is neither blue or red")
	}

	//switch

	switch colour {
	case "red":
		fmt.Println("colour is red")
	case "blue":
		fmt.Println("colour is blue")
	default:
		fmt.Println("colour is not blue or red")
	}

}
