package main

import (
	"fmt"
	"math"
)

func sincos() {
	var no, factorial int
	var isPlus bool
	var xDeg, xRad, term, sinx, cosx float64
	sinx = 0
	cosx = 0
	factorial = 1
	isPlus = true
	fmt.Printf("Enter the value of x (in degrees): ")
	fmt.Scanf("%f", &xDeg)
	fmt.Printf("\nEnter the number of terms: ")
	fmt.Scanf("%d", &no)
	xRad = xDeg * 3.14 / 180
	fmt.Printf("\n\nSine series = ")
	for i := 1; i < 2*no; i += 2 {
		factorial = 1
		for k := 1; k <= i; k++ {
			factorial *= k
		}
		term = (math.Pow(xRad, float64(i))) / float64(factorial)
		if isPlus {
			fmt.Printf("+")
			sinx += term
			isPlus = false
		} else {
			fmt.Printf("-")
			sinx -= term
			isPlus = true
		}
		fmt.Printf(" %f ", term)
	}
	fmt.Printf("\n\nsin %f = %f \n", xDeg, sinx)

	isPlus = true
	fmt.Printf("\n\nCosine series = ")
	for i := 0; i < 2*no; i += 2 {
		factorial = 1
		for k := 1; k <= i; k++ {
			factorial *= k
		}
		term = (math.Pow(xRad, float64(i))) / float64(factorial)
		if isPlus {
			fmt.Printf("+")
			cosx += term
			isPlus = false
		} else {
			fmt.Printf("-")
			cosx -= term
			isPlus = true
		}
		fmt.Printf(" %f ", term)
	}
	fmt.Printf("\n\ncos %f = %f \n", xDeg, cosx)
}
