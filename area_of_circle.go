package main

import (
	"fmt"
)

func main() {
	var r float64
	const pi float64 = 3.141592653
	var a float64
	fmt.Print("Enter Value Of Your Radius:\n")
	fmt.Scan(&r)
	a = pi * r * r
	fmt.Print("Area Will Be :\n", a)
}
