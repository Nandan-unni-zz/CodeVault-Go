package main

import (
	"fmt"
)

func area() {
	var lr, br, areaR, ht, bt, areaT float32
	fmt.Printf("Enter the length and breadth of Rectangle : \n")
	fmt.Scanf("%f", &lr)
	fmt.Scanf("%f", &br)
	areaR = lr * br
	fmt.Printf("Area of the Rectangle is %f \n", areaR)
	fmt.Printf("\nEnter the height and breadth of Triangle : \n")
	fmt.Scanf("%f", &ht)
	fmt.Scanf("%f", &bt)
	areaT = 0.5 * ht * bt
	fmt.Printf("Area of the Triangle is %f", areaT)
	return
}
