// this is go file
package main

import "fmt"

func main() {
	//Define map
	//emails := make(map[string]string)

	//Assign key values
	//emails["Bob"] = "bob@gmail.com"
	//emails["Travis"] = "travis@gmail.com"
	//emails["Mike"] = "mike@gmail.com"

	//declare map and add key values
	emails := map[string]string{"Bob": "bob@gmail.com", "travis": "travis@gmail.com"}

	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	//deleted one
	delete(emails, "Bob")

	//reprint (without bob)
	fmt.Println(emails)

}
