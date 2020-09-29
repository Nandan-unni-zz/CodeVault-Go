package main

import (
	"fmt"
)

type Student struct {
	name            string
	m1, m2, m3, tot int
	perc            float64
}

func main() {
	var n int
	var s [50]Student
	fmt.Printf("Enter the number of students : ")
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Printf("\nStudent %d : \nEnter Name : ", i+1)
		fmt.Scanf("%s", &s[i].name)
		fmt.Printf("Enter the marks in Subject 1 :")
		fmt.Scanf("%d", &s[i].m1)
		fmt.Printf("Enter the marks in Subject 2 :")
		fmt.Scanf("%d", &s[i].m2)
		fmt.Printf("Enter the marks in Subject 3 :")
		fmt.Scanf("%d", &s[i].m3)
		s[i].tot = s[i].m1 + s[i].m2 + s[i].m3
		s[i].perc = float64(s[i].tot / 3)
	}
	for i := 0; i < n; i++ {
		fmt.Printf("\n%s\n", s[i].name)
		fmt.Printf("Subject 1: %d\nSubject 2: %d\nSubject 3: %d\n", s[i].m1, s[i].m2, s[i].m3)
		fmt.Printf("Total marks: %d\nPercentage: %f\n", s[i].tot, s[i].perc)
	}
}
