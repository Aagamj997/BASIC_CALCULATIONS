package main

import (
	"fmt"
)

func main() {
	var a float64
	var b float64
	fmt.Print("Enter Your Length :\n")
	fmt.Scan(&a)
	b = a * a
	fmt.Print(" The Area Will Be :\n", b)
}
