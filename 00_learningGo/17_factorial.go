package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 1 {
		return n
	}
	return n * factorial(n-1)
}

func factMain() {
	var n int
	fmt.Printf("Enter the number : ")
	fmt.Scanf("%d", &n)
	fmt.Printf("%d! = %d\n", n, factorial(n))
}
