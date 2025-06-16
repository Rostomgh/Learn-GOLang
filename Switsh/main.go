package main

import (
	"fmt"
)

func main() {
	var a int = 20
	b := 14

	switch a < b {
	case true:
		fmt.Println("a is greater")
	case false:
		fmt.Println("b is greater")
	default:
		fmt.Println("a and b are equal")

	}

}
