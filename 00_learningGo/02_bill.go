package main

import (
	"fmt"
)

func bill() {
	var cn int
	var cu, rate float32
	fmt.Printf("Enter the Consumer no. : ")
	fmt.Scanf("%d", &cn)
	fmt.Printf("Enter the consumption units : ")
	fmt.Scanf("%f", &cu)
	if (cu >= 0) && (cu <= 200) {
		rate = 0.5 * cu
	} else if (cu >= 201) && (cu <= 400) {
		rate = 100 + (0.65 * (cu - 200))
	} else if (cu >= 401) && (cu <= 600) {
		rate = 230 + (0.8 * (cu - 400))
	} else {
		rate = 390 + (cu - 600)
	}
	fmt.Printf("\nRate of charge is %f", rate)
}
