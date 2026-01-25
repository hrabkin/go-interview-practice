package main

import (
	"fmt"
	"math"
)

func main() {
	// Example usage
	celsius := 25.0
	fahrenheit := CelsiusToFahrenheit(celsius)
	fmt.Printf("%.2f°C is equal to %.2f°F\n", celsius, fahrenheit)

	fahrenheit = 68.0
	celsius = FahrenheitToCelsius(fahrenheit)
	fmt.Printf("%.2f°F is equal to %.2f°C\n", fahrenheit, celsius)
}

// CelsiusToFahrenheit converts a temperature from Celsius to Fahrenheit
// Formula: F = C × 9/5 + 32
func CelsiusToFahrenheit(celsius float64) float64 {
	
	celsius = ValidateCelsius(celsius)
	res := celsius * 9/5 + 32
	return Round(res, 2)
}

// FahrenheitToCelsius converts a temperature from Fahrenheit to Celsius
// Formula: C = (F - 32) × 5/9
func FahrenheitToCelsius(fahrenheit float64) float64 {
    
    fahrenheit = ValidateFahrenheit(fahrenheit)
    res := (fahrenheit - 32) * 5/9
	return Round(res, 2)
}

func ValidateCelsius(celsius float64) float64 {
    if celsius < -273.15 {
        fmt.Printf("temperature below absolute zero: %f°C", celsius)
        return -273.15
    }
    return celsius
}

// Check for absolute zero violation in Fahrenheit
func ValidateFahrenheit(fahrenheit float64) float64 {
    if fahrenheit < -459.67 {
        fmt.Printf("temperature below absolute zero: %f°F", fahrenheit)
        return -459.67
    }
    return fahrenheit
}

// Round rounds a float64 value to the specified number of decimal places
func Round(value float64, decimals int) float64 {
	precision := math.Pow10(decimals)
	return math.Round(value*precision) / precision
}
