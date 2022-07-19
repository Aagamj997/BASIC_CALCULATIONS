package main

import (
	"fmt"
)

func divide(a int, b int) int {
	var c = a / b
	return c
}
func main() {
	var a, b, c int
	fmt.Print("Enter Your Dividend :\n")
	fmt.Scan(&a)
	fmt.Print("Enter Your Divisor :\n")
	fmt.Scan(&b)
	c = divide(a, b)
	fmt.Printf("Your Quotient is :%d", c)
}
