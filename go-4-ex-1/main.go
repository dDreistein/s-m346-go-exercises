package main

import "fmt"

func computeGrade(gotPoints float32, maxPoints float32) float32 {
	return (gotPoints/maxPoints)*5 + 1
}

func main() {
	fmt.Println(computeGrade(1.0, 2.0))       // 3.5
	fmt.Println(computeGrade(134.123, 167.0)) // 5.015659
	fmt.Println(computeGrade(74.0, 100.0))    // 4.7
}
