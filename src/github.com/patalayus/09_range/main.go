// this is go file
package main

import "fmt"

func main() {
	ids := []int{33, 76, 34, 63, 45, 23, 3}

	//loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//not using index
	//use underscore if variable is unused
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	//add IDs together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	//this is the sum of IDs
	fmt.Println("\nSum", sum)

	//Range with map
	emails := map[string]string{"Bob": "bob@gmail.com", "Travis": "travis@gmail.com"}

	//using key value pairs
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}
}
