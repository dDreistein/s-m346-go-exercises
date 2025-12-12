package main

import (
	"fmt"
	"math"
)

func computeHypotenuse(a float64, b float64) float64 {
	return math.Sqrt(a*a + b*b) // Ist übersichtlicher als math.Pow()
}

func main() {
	fmt.Println(computeHypotenuse(1.0, 2.0))       // 2.23606797749979
	fmt.Println(computeHypotenuse(5, 12.0))        // 13
	fmt.Println(computeHypotenuse(3.2345, 4.6789)) // 5.688066056226844
}
