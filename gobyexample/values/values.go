package main

import "fmt"

func main() {
	// Concatenate both strings
	fmt.Println("go" + "lang") // Output: golang

	// Spaces after every value
	fmt.Println("1+1 =", 1+1)         // Output: 1+1 = 2
	fmt.Println("7.0/3.0 =", 7.0/3.0) // Output: 7.0/3.0 = 2.333...

	// Comparison
	fmt.Println(true && false) // Output: false
	fmt.Println(true || false) // Output: true
	fmt.Println(!true)         // Output: false
}
