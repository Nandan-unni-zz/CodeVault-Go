package main

import (
	"fmt"
)

func matrix() {
	var a, b, sum, product [20][20]int
	var rowA, colA, rowB, colB int
	fmt.Printf("\nEnter the rows and column of First Matrix : ")
	fmt.Scanf("%d", &rowA)
	fmt.Scanf("%d", &colA)
	fmt.Printf("\nEnter the rows and column of Second Matrix : ")
	fmt.Scanf("%d", &rowB)
	fmt.Scanf("%d", &colB)
	if colA != rowB {
		fmt.Printf("\nMatrix Addition/Multiplication not possible.\n")
		return
	}
	fmt.Printf("\nEnter the Elements of Matrix A : \n")
	for i := 0; i < rowA; i++ {
		for j := 0; j < colA; j++ {
			fmt.Scanf("%4d", &a[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\nEnter the Elements of Matrix B : \n")
	for i := 0; i < rowB; i++ {
		for j := 0; j < colB; j++ {
			fmt.Scanf("%4d", &b[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\nMatrix A: \n")
	for i := 0; i < rowA; i++ {
		for j := 0; j < colA; j++ {
			fmt.Printf("%4d", a[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\nMatrix B: \n")
	for i := 0; i < rowB; i++ {
		for j := 0; j < colB; j++ {
			fmt.Printf("%4d", b[i][j])
		}
		fmt.Printf("\n")
	}
	if rowA != rowB || colA != colB {
		fmt.Printf("\nMatrix addition not possible !")
	} else {
		fmt.Printf("\nSum of Matrix A and Matrix B :\n")
		for i := 0; i < rowA; i++ {
			for j := 0; j < colA; j++ {
				sum[i][j] = a[i][j] + b[i][j]
				fmt.Printf("%4d", sum[i][j])
			}
			fmt.Printf("\n")
		}
	}
	fmt.Printf("\nProduct of Matrix A and Matrix B :\n")
	for i := 0; i < rowA; i++ {
		for j := 0; j < colB; j++ {
			product[i][j] = 0
			for k := 0; k < colA; k++ {
				product[i][j] += a[i][k] * b[k][j]
			}
			fmt.Printf("%4d", product[i][j])
		}
		fmt.Printf("\n")
	}
}
