package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&b)
	a = 14

	if a > b {
		fmt.Println("a is greater")
	} else if a < b {
		fmt.Println("b is greater")
	} else {
		fmt.Println("a and b are equal")
	}

}
