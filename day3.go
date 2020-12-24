package main

import "fmt"

func main() {
	var (
		x = 5
		y = 15
	)
	fmt.Println(x > y)
	fmt.Println(x < y)
	fmt.Println(x >= y)
	fmt.Println(x <= y)
	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x > y && x == y)
	fmt.Println(!(x <= y))
	fmt.Println(x == y || x < y)
}
