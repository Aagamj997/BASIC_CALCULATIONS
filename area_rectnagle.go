package main

import (
	"fmt"
)

func main() {
	var l float64
	var b float64
	var a float64
	fmt.Print("Enter Your Length:")
	fmt.Scan(&l)
	fmt.Print("Enter Your Breadth :")
	fmt.Scan(&b)
	a = l * b
	fmt.Print("The Area Will Be: ", a)
}
