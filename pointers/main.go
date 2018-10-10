package main

import (
	"fmt"
)

func main() {
	a := 10
	b := &a
	fmt.Println(a, b)
	fmt.Printf("%T\n", b) //*int
	// Use * to read address
	fmt.Println("address's value:", *b)
	*b = 5
	fmt.Println(a)
}
