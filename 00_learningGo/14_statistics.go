package main

import (
	"fmt"
	"math"
)

func mean(data [15]int, n int) float64 {
	sum := 0
	var mn float64
	for i := 0; i < n; i++ {
		sum += data[i]
	}
	mn = float64(sum / n)
	return mn
}

func varience(data [15]int, n int) float64 {
	var sqSum, vr float64
	sqSum = 0
	for i := 0; i < n; i++ {
		sqSum += float64(data[i] * data[i])
	}
	vr = sqSum / float64(n)
	return vr
}

func sdv(data [15]int, n int) float64 {
	vr := varience(data, n)
	sd := math.Sqrt(vr)
	return sd
}

func statistics() {
	var data [15]int
	var n int
	fmt.Printf("Enter the no of elements : ")
	fmt.Scanf("%d", &n)
	fmt.Printf("\nEnter the dataset : \n")
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &data[i])
	}
	fmt.Printf("The of mean the data set is %f\n", mean(data, n))
	fmt.Printf("The of varience the data set is %f\n", varience(data, n))
	fmt.Printf("The of standard deviation the data set is %f\n", sdv(data, n))
}
