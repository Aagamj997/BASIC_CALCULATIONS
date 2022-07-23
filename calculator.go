package main

import (
	"fmt"
)

func main() {
	var a int
	var x float64
	var y float64
	fmt.Println("WHAT DO YOU WANT TO DO ? \n 1.Addition \n 2.Subtraction \n 3. Multiplication \n 4.Division")
	fmt.Scan(&a)
	if a == 1 {
		fmt.Print("Enter Your First Number :\n")
		fmt.Scan(&x)
		fmt.Print("Enter Your Next Number :\n")
		fmt.Scan(&y)
		fmt.Print("Your Sum Will Be \n ", (x + y))
	}
	if a == 2 {
		fmt.Print("Enter Your First Number :\n")
		fmt.Scan(&x)
		fmt.Print("Enter Your Next Number :\n")
		fmt.Scan(&y)
		fmt.Print("Your Diff Will Be \n ", (x - y))
	}
	if a == 3 {
		fmt.Print("Enter Your First Number :\n")
		fmt.Scan(&x)
		fmt.Print("Enter Your Next Number :\n")
		fmt.Scan(&y)
		fmt.Print("Your Product Will Be \n ", (x * y))
	}
	if a == 4 {
		fmt.Print("Enter Your First Number :\n")
		fmt.Scan(&x)
		fmt.Print("Enter Your Next Number :\n")
		fmt.Scan(&y)
		fmt.Print("Your Quotient Will Be \n ", (x / y))
	}
	if a > 4 {
		fmt.Print("Invalid Input")
	}
}
