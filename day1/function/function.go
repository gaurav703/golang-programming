package main

import "fmt"

func concat(a string, b string) string {
	return a + " " + b
}

func main() {
	fmt.Println("Hello, World!", concat("Hello", "World"))

	// function with multiple return values
	a, b := 10, 20
	sum, diff := addSubtract(a, b)
	fmt.Println("Sum is", sum)
	fmt.Println("Difference is", diff)

	// ignore return values
	_, diff = addSubtract(a, b)
	fmt.Println("Difference is ignore value ===",a , "-", b , "==" , diff)

}

func addSubtract(a int, b int) (int, int) {
	return a + b, a - b
}

// In the above code, we have defined a function concat that takes two string arguments and returns a string. We have also defined a function addSubtract that takes two integer arguments and returns two integers. We have called these functions in the main function. The addSubtract function returns two values, and we have used both values in the main function. We have also ignored one of the return values in the main function.
//
// Run this code using the following command:
// go run function.go
//
