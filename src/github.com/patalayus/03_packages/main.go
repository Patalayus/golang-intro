// this is go file
package main

import (
	"fmt" 
	"math"
	"github.com/patalayus/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	//importing custom packages
	fmt.Println(strutil.Reverse("hello"))

}
