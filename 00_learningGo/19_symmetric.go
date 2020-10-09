// A S Nandanunni
// Reg No: 20219023
// CS - A

package main

import "fmt"

func main() {
	var a [10][10]int
	var row, col, choice int
	fmt.Printf("\nEnter the rows and column of the Matrix : ")
	fmt.Scanf("%d %d", &row, &col)
	fmt.Printf("\nEnter the Elements of the Matrix: \n")
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Scanf("%4d", &a[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("1. Check whether matrix is symmetric")
	fmt.Printf("\n2. Display the upper and lower matrix")
	fmt.Printf("\nEnter your choice : ")
	fmt.Scanf("%d", &choice)
	switch choice {
	case 1:
		isSymmetric(a, row, col)
		break
	case 2:
		upper(a, row, col)
		lower(a, row, col)
		break
	default:
		fmt.Printf("\nInvalid choice !")
		break
	}
}

func isSymmetric(a [10][10]int, row int, col int) {
	var transpose [10][10]int
	var result int
	result = 1
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			transpose[j][i] = a[i][j]
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if transpose[i][j] != a[i][j] {
				result = 0
			}
		}
	}
	if result == 1 {
		fmt.Printf("The array is symmetric")
	} else {
		fmt.Printf("The array is not symmetric")
	}
}

func upper(a [10][10]int, row int, col int) {
	fmt.Printf("\n\n The upper matrix is : ")
	for i := 0; i < row; i++ {
		fmt.Printf("\n")
		for j := 0; j < col; j++ {
			if i > j {
				fmt.Printf("0")
				fmt.Printf("\t")
			} else {
				fmt.Printf("%d\t", a[i][j])
			}
		}
	}
}

func lower(a [10][10]int, row int, col int) {
	fmt.Printf("\n\n The lower matrix is : ")
	for i := 0; i < row; i++ {
		fmt.Printf("\n")
		for j := 0; j < col; j++ {
			if i >= j {
				fmt.Printf("%d\t", a[i][j])
			} else {
				fmt.Printf("0")
				fmt.Printf("\t")
			}
		}
	}
}
