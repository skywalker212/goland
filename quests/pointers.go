package main

import "fmt"

func main() {
	x := 15
	a := &x // memory address
	fmt.Println("Address:", a)
	fmt.Println("Value:", *a)
	*a = *a * *a
	fmt.Println(*a, x)
}
