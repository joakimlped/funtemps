package main

import (
	"flag"
	"fmt"
	"github.com/joakimlped/funtemps/conv"
)

var (
	fahrenheit float64
	celsius    float64
	kelvin     float64
	outUnit    string
)

func init() {
	flag.Float64Var(&fahrenheit, "F", 0.0, "temperature in degrees Fahrenheit")
	flag.Float64Var(&celsius, "C", 0.0, "temperature in degrees Celsius")
	flag.Float64Var(&kelvin, "K", 0.0, "temperature in Kelvin")
	flag.StringVar(&outUnit, "out", "C", "output unit, can be 'C' for Celsius, 'F' for Fahrenheit or 'K' for Kelvin")
}

func main() {
	flag.Parse()

	// Check that only one input flag is set
	numInputFlags := 0
	for _, arg := range []float64{fahrenheit, celsius, kelvin} {
		if arg != 0 {
			numInputFlags++
		}
	}
	if numInputFlags != 1 {
		fmt.Println("Error: please specify one input flag only (-C, -F, -K)")
		return
	}

	var inputTemp, outputTemp float64
	var inputUnit string

	if fahrenheit != 0 {
		inputTemp = fahrenheit
		inputUnit = "F"
	} else if celsius != 0 {
		inputTemp = celsius
		inputUnit = "C"
	} else if kelvin != 0 {
		inputTemp = kelvin
		inputUnit = "K"
	}

	switch outUnit {
	case "C":
		if inputUnit == "F" {
			outputTemp = conv.FahrenheitToCelsius(inputTemp)
		} else if inputUnit == "K" {
			outputTemp = conv.KelvinToCelsius(inputTemp)
		} else {
			outputTemp = inputTemp
		}
		fmt.Printf("%.2f%s = %.2fC\n", inputTemp, inputUnit, outputTemp)

	case "F":
		if inputUnit == "C" {
			outputTemp = conv.CelsiusToFahrenheit(inputTemp)
		} else if inputUnit == "K" {
			outputTemp = conv.KelvinToFahrenheit(inputTemp)
		} else {
			outputTemp = inputTemp
		}
		fmt.Printf("%.2f%s = %.2fF\n", inputTemp, inputUnit, outputTemp)

	case "K":
		if inputUnit == "C" {
			outputTemp = conv.CelsiusToKelvin(inputTemp)
		} else if inputUnit == "F" {
			outputTemp = conv.FahrenheitToKelvin(inputTemp)
		} else {
			outputTemp = inputTemp
		}
		fmt.Printf("%.2f%s = %.2fK\n", inputTemp, inputUnit, outputTemp)

	default:
		fmt.Printf("Error: invalid output unit specified: %s\n", outUnit)
		return
	}
}
