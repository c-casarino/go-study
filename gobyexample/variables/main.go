package main

import "fmt"

func main() {
	// var a = "initial" -> type will be inferred from the rigth-hand side
	a := "initial" // Type: string
	fmt.Println(a) // Output: initial

	b, c := 1, 2 // Type: int
	fmt.Println(b, c)

	d := true      // Type = bool
	fmt.Println(d) // Output: true

	var e int      // Default: 0
	fmt.Println(e) // Output: 0 --> because zero is the default value of int
}
