package main

import "fmt"

func main() {
	celsius := 36.50
	fmt.Println(convertTemperature(celsius))

	//Note that:

	//Kelvin = Celsius + 273.15
	//Fahrenheit = Celsius * 1.80 + 32.00
}
func convertTemperature(celsius float64) []float64 {
	var result []float64
	kelvin := celsius + 273.15000
	Fahrenheit := celsius*1.80 + 32.00000
	result = append(result, kelvin+0.000, Fahrenheit+0.000)

	return result
}
