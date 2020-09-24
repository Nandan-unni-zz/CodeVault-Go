package main

import (
	"fmt"
)

func avg() {
	var no, sum int
	var a [20]int
	var avg float32
	sum = 0
	fmt.Printf("\n Enter the number of elements in the array: ")
	fmt.Scanf("%d", &no)
	fmt.Printf("\n Enter the elements: \n")
	for i := 0; i < no; i++ {
		fmt.Scanf("%d", &a[i])
		sum += a[i]
	}
	avg = float32(sum / no)
	fmt.Printf("\nThe sum of the elements in the given array is %d", sum)
	fmt.Printf("\nThe average of the elements in the given array is %f", avg)
}
