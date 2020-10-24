// this is go file
package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	// define Person struct
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " +
		strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// init person using struct
	// person1 := Person {firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	// fmt.Println(person1)

	// Alternative declaration

	person1 := Person{"Samantha", "Smith", "Boston", "f", 25}
	fmt.Println(person1)

	person2 := Person{"Bob", "Johnson", "Coventry", "m", 34}

	// get single field
	//fmt.Println(person1.firstName)
	//person1.age++
	//fmt.Println(person1)

	person1.hasBirthday()
	person1.hasBirthday() //adds birthday (value reciever)
	person1.hasBirthday() //adds birthday again

	person1.getMarried("Williams") //gets married
	person2.getMarried("Thomson")  //bob gets married (doesnt change name)

	fmt.Println(person2.greet())
	fmt.Println(person1.greet())

}
