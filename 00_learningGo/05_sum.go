package main

import (
	"fmt"
)

func sum() {
	var num, copy, rem, sum int
	fmt.Printf("Enter a number : ")
	fmt.Scanf("%d", &num)
	copy = num
	sum = 0
	for copy > 0 {
		rem = copy % 10
		sum += rem
		copy /= 10
	}
	fmt.Printf("\nThe sum of the digits of the number %d is %d", num, sum)
}
