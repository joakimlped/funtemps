package conv

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

    // Temperatur konvertering
    var convertedTemp float64
    switch *output {
    case "F":
        convertedTemp = CelsiusToFahrenheit(KelvinToCelsius(*input))
    case "K":
        convertedTemp = CelsiusToKelvin(KelvinToCelsius(*input))
    default:
        convertedTemp = KelvinToCelsius(*input)
    }

    // Output temperatur resultat
    fmt.Printf("%g%s is %g%s\n", *input, *output, convertedTemp, *output)

    // Print fun facts if requested, ikke fungerende enn så lenge
    if *funFacts != "" {
        switch *funFacts {
        case "sun":
            fmt.Println("Temperatur på ytre lag av Solen er 5506.85°C.")
            fmt.Println("Temperatur i Solens kjerne er 15,000,000°C.")
        case "moon", "luna":
            fmt.Println("Temperatur på Månens overflate om dagen 107°C.")
            fmt.Println("Temperatur på Månens overflate om natten er -153°C.")
        case "earth", "terra":
            fmt.Println("Laveste temperatur målt på Jordens overflate er -89.2°C.")
            fmt.Println("Høyeste temperatur målt på Jordens overflate er 56.7°C.")
        default:
            fmt.Println("Uvisst funfact er forespurt")
        }
    }
}
