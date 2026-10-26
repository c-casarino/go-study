package main

import (
	"fmt"
	"log"
)

func sum(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func multi(a, b float64) float64 {
	return a * b
}

func div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}

	return a / b, nil
}

func main() {
	resultSum := sum(5, 10)
	resultSub := sub(10, 5)
	resultMulti := multi(10, 5)
	resultDiv, err := div(12, 2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result Sum: ", resultSum)
	fmt.Println("Result Sub: ", resultSub)
	fmt.Println("Result Multi: ", resultMulti)
	fmt.Println("Result Div: ", resultDiv)
}
