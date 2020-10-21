// this is go file
package main

import "fmt"

func main() {
	//Arrays
	//var fruitArr [2]string

	//Assign values
	//fruitArr[0] = "Apple"
	//fruitArr[1] = "Orange"

	//Declare and assign
	//fruitArr := [2]string{"Apple", "Orange"} //shorthand

	//fmt.Println(fruitArr)
	//fmt.Println(fruitArr[1]) //specify item
	fruitSlice := [3]string{"Apple", "Orange", "Grape", "Cherry"}
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3]) //starts at 1 and stops before 3

}
