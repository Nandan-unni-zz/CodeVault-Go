package main

import (
	"fmt"
)

func prime() {
	var limit int
	var isPrime bool
	fmt.Printf("Enter the limit : ")
	fmt.Scanf("%d", &limit)
	if limit > 2 {
		for i := 2; i <= limit; i++ {
			isPrime = true
			for j := 2; j <= i/2; j++ {
				if i%j == 0 {
					isPrime = false
				}
			}
			if isPrime {
				fmt.Printf("%d, ", i)
			}
		}
	} else {
		fmt.Printf("only 2")
	}
}
