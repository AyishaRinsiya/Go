package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	var y float64
	var z float64
	var aa float64
	var c float64

	fmt.Print("Enter the value of x:")
	fmt.Scan(&x)
	fmt.Print("Enter the value of y:")
	fmt.Scan(&y)
	fmt.Print("Enter the value of z:")
	fmt.Scan(&z)
	fmt.Print("Enter the value of aa:")
	fmt.Scan(&aa)

	c = ((math.Sqrt(x) + math.Sqrt(y)*z) / aa)

	fmt.Println("(root(x) + root(y)*z)/aa)= ", c)

}
