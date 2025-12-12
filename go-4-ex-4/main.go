package main

import "fmt"

func convertCelsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func convertFahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {
	fmt.Println(convertCelsiusToFahrenheit(273.15)) // 523.67
	fmt.Println(convertCelsiusToFahrenheit(0))      // 32
	fmt.Println(convertCelsiusToFahrenheit(100))    // 212

	fmt.Println(convertFahrenheitToCelsius(32))     // 0
	fmt.Println(convertFahrenheitToCelsius(212))    // 100
	fmt.Println(convertFahrenheitToCelsius(491.67)) // 255.3722222222222

}
