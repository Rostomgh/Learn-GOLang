package main

import (
	"fmt"
	"strings"
)

func main() {
	//Compare
	A := "hello"
	B := "world"
	// fmt.Println(strings.Compare(A, B))
	//Contains
	fmt.Println(strings.Contains(A, "o"))
	fmt.Println(strings.Contains(B, "W"))
	fmt.Println(strings.ContainsAny(B, "wwworledddds"))

}
