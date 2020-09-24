package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, e, d, e string
	la := 0
	lb := 0
	lc := 0
	ld := 0
	isSame := true
    fmt.Printf("Enter string A : ");
    fmt.Scanf("%s", a);
    for i := 0; a[i] != '\0'; i++ {
		la++;
	}
    lb = la;
    lc = la;
    ld = lb + la;
    for i, j := 0, lb - 1; i < la; i, j = i+1, j+1 {
		b[j] = a[i]
	}
    for i := 0; i < la; i++ {
        c[i] = a[i]
        d[i] = a[i]
	}
	for i, ch := range a {
		c[i] = a[i]
	}
    for i, j := la, 0; i < ld; i, j = i+1, j+1 {
		d[i] = b[j]
	}
    fmt.Printf("String A : %s\n", a);
    fmt.Printf("Length of String A : %d\n", la);
    fmt.Printf("String B (Reversed A) : %s\n", b);
    fmt.Printf("String C (Copied String) : %s\n", c);
    fmt.Printf("String D (Concatenated String) : %s\n", d);
    fmt.Printf("Enter String E of %d length to compare with string A : ", la);
    fmt.Scanf("%s", e);
    for i := 0; i < la; i++ {
        if (a[i] != e[i]) {
			isSame = false;
		}
	}
    if (isSame) {
		fmt.Printf("String A and String E are same.");
	} else {
		fmt.Printf("String A and String E are different.\n");
	}
}
