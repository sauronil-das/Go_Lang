package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y float64
	fmt.Printf("Enter First Floating Point Value: ")
	fmt.Scan(&x)
	x = math.Trunc(x)
	fmt.Printf("\nTruncated Value: %.0f", x)

	fmt.Printf("\nEnter Second Floating Point Value: ")
	fmt.Scan(&y)
	y = math.Trunc(y)
	fmt.Printf("\nTruncated Value: %.0f\n", y)
}
