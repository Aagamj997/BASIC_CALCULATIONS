package main

import (
	"fmt"
)

func main() {
	var a float64
	var b float64
	var c float64
	fmt.Print("Enter Length Of First Diagonal Of Rhombus :\n")
	fmt.Scan(&a)
	fmt.Print("Enter Your Length Of Second Diuagobal :\n")
	fmt.Scan(&b)
	c = ((0.5) * a * b)
	fmt.Print("The Area Of Rhombus Will Be ", c)
}
