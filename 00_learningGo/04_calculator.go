package main

import (
	"fmt"
	"math"
)

func calculator() {
	var a, b, c float32
	var op string
	var isValid bool = true
	fmt.Print("Enter the first number : ")
	fmt.Scanln(&a)
	fmt.Print("Enter the operator : ")
	fmt.Scanln(&op)
	fmt.Print("Enter the first number : ")
	fmt.Scanln(&b)
	switch op {
	case "+":
		c = a + b
	case "-":
		c = a - b
	case "*":
		c = a * b
	case "/":
		if b != 0 {
			c = a / b
		} else {
			fmt.Print("\nDivision by 0 not possible !\n")
			isValid = false
		}
	case "^":
		c = float32(math.Pow(float64(a), float64(b)))
	case "%":
		c = float32(int(a) % int(b))
	default:
		fmt.Print("\nInvalid Operator !\n")
		isValid = false
	}
	if isValid {
		fmt.Printf("\n%g %s %g = %g\n", a, op, b, c)
	}
}
