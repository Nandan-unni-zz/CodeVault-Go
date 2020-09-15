package main

import (
	"fmt"
)

func main() {
	var a, b string
	var c = []string{"a", "b", "c"}
	d := []int{1, 2, 3}
	fmt.Print("Enter string a : ")
	fmt.Scan(&a)
	fmt.Printf("Enter string b : ")
	fmt.Scanf("%s", &b)
	print(a)
	fmt.Print("Print : To print simple text.\n")
	fmt.Printf("Printf : Prints formatted strings like a = %s and b = %s \n", a, b)
	for i, char := range c {
		fmt.Printf("%d --> %s\n", i, char)
	}
	for true {
		for i := range d {
			fmt.Printf("While loop : %d\n", i)
		}
		break
	}
}
