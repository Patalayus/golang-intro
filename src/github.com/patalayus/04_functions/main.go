// this is go file
package main

import "fmt" //formatter module

func greeting(name string) string{
	return "Hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("alex"))
	fmt.Println(getSum(3, 4))
}
