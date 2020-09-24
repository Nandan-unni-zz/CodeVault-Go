package main

import (
	"fmt"
	"math"
)

func roots() {
	var a, b, c, d, e, f float64
	fmt.Printf("The quadratic eqn is a(x*x) + bx + c = 0")
	fmt.Printf("\nEnter the coefficients :\n  a: ")
	fmt.Scanf("%f", &a)
	fmt.Printf("  b: ")
	fmt.Scanf("%f", &b)
	fmt.Printf("  c: ")
	fmt.Scanf("%f", &c)
	fmt.Printf("\nThe quadratic eqn is %f(x*x) + %fx + %f = 0", a, b, c)
	f = (b * b) - (4 * a * c)
	if f >= 0 {
		d = (b + math.Sqrt(f)) / (2 * a)
		e = (b - math.Sqrt(f)) / (2 * a)
		fmt.Printf("\nThe roots of the quadratic eqn are %f and %f", d, e)
	} else {
		fmt.Printf("\nImaginary roots..!")
	}
}
