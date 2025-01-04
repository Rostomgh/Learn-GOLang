package main

import (
	"fmt"
)

func main() {
	//Scanner input
	fmt.Printf("what is your name?\n")
	var name string
	fmt.Scanln(&name)
	fmt.Print("my name is ", name)
}
