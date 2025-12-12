package main

import "math"
import "fmt"

// TODO: implement the function computeQuadraticFormula
func computeQuadraticFormula(a, b, c float64) []float64 {
	discriminant := b*b - 4*a*c

	var solutions []float64
	if discriminant > 0 {
		solution1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		solution2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		solutions = append(solutions, solution1, solution2)
	} else if discriminant == 0 {
		solution := -b / (2 * a)
		solutions = append(solutions, solution)
	}
	return solutions
}

func main() {
	// TODO: call the function computeQuadraticFormula
	fmt.Println(computeQuadraticFormula(3, 4, 1)) // [-0.3333333333333333 -1]
	fmt.Println(computeQuadraticFormula(2, 4, 2)) // [-1]
	fmt.Println(computeQuadraticFormula(3, 4, 2)) // []
}
