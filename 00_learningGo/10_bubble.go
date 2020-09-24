package main

import (
	"fmt"
)

func bubble() {
	var no, temp int
	var a [20]int
	fmt.Printf("\nEnter the number of elements in the array: ")
	fmt.Scanf("%d", &no)
	fmt.Printf("\nEnter the elements in the array: \n")
	for i := 0; i < no; i++ {
		fmt.Scanf("%d", &a[i])
	}
	fmt.Printf("\nThe Input array is : [ ")
	for i := 0; i < no; i++ {
		fmt.Printf("%d, ", a[i])
	}
	fmt.Printf("]\n")
	for i := 0; i < no; i++ {
		for j := 0; j < no-i-1; j++ {
			if a[j] > a[j+1] {
				temp = a[j]
				a[j] = a[j+1]
				a[j+1] = temp
			}
		}
	}
	fmt.Printf("\nThe Sorted Array is : [ ")
	for i := 0; i < no; i++ {
		fmt.Printf("%d, ", a[i])
	}
	fmt.Printf("]\n")
}
