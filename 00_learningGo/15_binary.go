package main

import (
	"fmt"
)

func linear(a [20]int, n int, el int) {
	found := false
	for i := 0; i < n; i++ {
		if a[i] == el {
			fmt.Printf("%d found at position %d\n", el, i+1)
			found = true
		}
	}
	if !found {
		fmt.Printf("%d not found in array\n", el)
	}
}
func binary(a [20]int, n int, el int) {
	var start, mid, end int
	var found bool
	start = 0
	mid = n / 2
	end = n - 1
	for start < end {
		if el == a[mid] {
			fmt.Printf("%d found at position %d\n", el, mid+1)
			found = true
			break
		} else if el < a[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	if !found {
		fmt.Printf("%d not found in arrray\n", el)
	}
}

func linBinary() {
	var n, ls, bs int
	var a [20]int
	fmt.Printf("Enter the no of elements : ")
	fmt.Scanf("%d", &n)
	fmt.Printf("Enter the elements in increasing order : \n")
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	fmt.Printf("\nLinear Search\nEnter the number to search : ")
	fmt.Scanf("%d", &ls)
	linear(a, n, ls)
	fmt.Printf("\nBinary Search\nEnter the number to search : ")
	fmt.Scanf("%d", &bs)
	binary(a, n, bs)
}
