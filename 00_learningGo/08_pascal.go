package main

import (
	"fmt"
)

func pascal() {
	var rows, no int
	fmt.Printf("\n Enter the number of rows: ")
	fmt.Scanf("%d", &rows)
	fmt.Printf("\n")
	for i := 0; i < rows; i++ {
		for k := 1; k < rows-i; k++ {
			fmt.Printf("  ")
		}
		for j := 0; j <= i; j++ {
			if j == 0 {
				no = 1
			} else {
				no = no * (i - j + 1) / j
			}
			fmt.Printf(" %2d ", no)
		}
		fmt.Printf("\n")
	}
}
