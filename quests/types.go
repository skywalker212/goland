package main

import "fmt"

// a simple function, mention the return type also if returning anyting
func add(x float32, y float32) float32 {
	return x + y
}

/*
 * A function that returns multiple things
 * you can club similar types by just mentioning the type at the end, in this case (a, b string)
 */
func multiple(a, b string) (string, string) {
	return a, b
}

/*
 * this is a multiple line comment
 * commenting out unused variables otherwise it will throw an error at compile time
 */
func main() {
	// gives float64 by default, be careful
	// num1, num2 := 5.5, 5.5

	// will create two string type variables
	w1, w2 := "hello", "there"
	fmt.Println(multiple(w1, w2))

	// convert from one to another
	// var a int = 32
	// var b float64 = float64(a)
	// x := a // x will be of type int
}
