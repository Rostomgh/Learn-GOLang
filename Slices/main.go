package main

import (
	"fmt"
)

func main() {
	myslice := []string{"a", "b", "c"}
	fmt.Println(myslice)
	myslice[1] = "d"
	fmt.Println(len(myslice))
	fmt.Println(myslice[1])
	fmt.Println(cap(myslice))
	fmt.Println(myslice)
}
