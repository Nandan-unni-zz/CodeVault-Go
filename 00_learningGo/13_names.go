package main

import (
	"fmt"
)

func name() {
	var names [20]string
	var temp string
	var n int
	fmt.Printf("Enter the number of names : ")
	fmt.Scanf("%d", &n)
	fmt.Printf("Enter the %d names : \n", n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &names[i])
	}
	for i := 0; i < n; i++ {
		k := 0
		for j := i; j < n; j++ {
			if names[i][k] > names[j][k] {
				temp = names[i]
				names[i] = names[j]
				names[j] = temp
			}
		}
	}
	fmt.Printf("\nThe names after sorting : \n")
	for i := 0; i < n; i++ {
		fmt.Printf("%s\n", names[i])
	}
}
