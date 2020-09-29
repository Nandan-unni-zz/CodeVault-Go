package main

import (
	"fmt"
)

func fibonacci(limit int) int {
	if limit <= 1 {
		return limit
	}
	return fibonacci(limit-1) + fibonacci(limit-2)

}

func fibMain() {
	var limit int
	fmt.Printf("Enter the limit of the series : ")
	fmt.Scanf("%d", &limit)
	for i := 0; i <= limit; i++ {
		fmt.Printf("%d, ", fibonacci(i))
	}
}
