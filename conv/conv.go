package main

import (
    "flag"
    "fmt"
)

func FahrenheitToCelsius(value float64) float64 {
    return (value - 32) * 5 / 9
}

func CelsiusToFahrenheit(value float64) float64 {
    return value*9/5 + 32
}

func KelvinToCelsius(value float64) float64 {
    return value - 273.15
}

func CelsiusToKelvin(value float64) float64 {
    return value + 273.15
}

func KelvinToFahrenheit(value float64) float64 {
    return (value-273.15)*9/5 + 32
}

func FahrenheitToKelvin(value float64) float64 {
    return (value-32)*5/9 + 273.15
}

func main() {
    input := flag.Float64("i", 0, "Input temperature value")
    output := flag.String("out", "C", "Output temperature unit (C, F, K)")
    funFacts := flag.String("funfacts", "", "Print fun facts about the Sun, Moon, or Earth")
    flag.Parse()

    // Temperature conversion
    var convertedTemp float64
    switch *output {
    case "F":
        convertedTemp = CelsiusToFahrenheit(KelvinToCelsius(*input))
    case "K":
        convertedTemp = CelsiusToKelvin(KelvinToCelsius(*input))
    default:
        convertedTemp = KelvinToCelsius(*input)
    }

    // Output temperature conversion result
    fmt.Printf("%g%s is %g%s\n", *input, *output, convertedTemp, *output)

    // Print fun facts if requested
    if *funFacts != "" {
        switch *funFacts {
        case "sun":
            fmt.Println("Temperature on the outer layer of the Sun is 5506.85°C.")
            fmt.Println("Temperature in the core of the Sun is 15,000,000°C.")
        case "moon", "luna":
            fmt.Println("The average temperature on the Moon during the day is 107°C.")
            fmt.Println("The average temperature on the Moon during the night is -153°C.")
        case "earth", "terra":
            fmt.Println("The coldest temperature ever recorded on Earth was -89.2°C.")
            fmt.Println("The hottest temperature ever recorded on Earth was 56.7°C.")
        default:
            fmt.Println("Unknown fun fact requested.")
        }
    }
}
