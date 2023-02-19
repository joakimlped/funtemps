package main

import (
	"flag"
	"fmt"
	"strings"
)

var fahr float64
var out string
var funfactsArg string
var tempScale string

func init() {
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfactsArg, "funfacts", "", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	flag.StringVar(&tempScale, "t", "C", "temperaturskala for fun-fact (C, F eller K)")
}

func main() {
	flag.Parse()

	if flag.NFlag() == 0 || (flag.NFlag() == 1 && isFlagPassed("funfacts")) {
		flag.Usage()
		return
	}

	if isFlagPassed("F") && isFlagPassed("C") {
		fmt.Println("Kan ikke bruke flaggene -F og -C samtidig.")
		return
	}

	if isFlagPassed("F") && isFlagPassed("K") {
		fmt.Println("Kan ikke bruke flaggene -F og -K samtidig.")
		return
	}

	if isFlagPassed("C") && isFlagPassed("K") {
		fmt.Println("Kan ikke bruke flaggene -C og -K samtidig.")
		return
	}

	if isFlagPassed("funfacts") && !isFlagPassed("t") {
		fmt.Println("Må spesifisere -t flagget når man bruker -funfacts.")
		return
	}

	if isFlagPassed("F") {
		if out == "C" {
			fmt.Printf("%g°F er %g°C\n", fahr, conv.FahrenheitToCelsius(fahr))
		} else if out == "K" {
			fmt.Printf("%g°F er %gK\n", fahr, conv.FahrenheitToKelvin(fahr))
		} else if out == "F" {
			fmt.Printf("%g°F\n", fahr)
		}
	} else if isFlagPassed("C") {
		if out == "F" {
			fmt.Printf("%g°C er %g°F\n", fahr, conv.CelsiusToFahrenheit(fahr))
		} else if out == "K" {
			fmt.Printf("%g°C er %gK\n", fahr, conv.CelsiusToKelvin(fahr))
		} else if out == "C" {
			fmt.Printf("%g°C\n", fahr)
		}
	} else if isFlagPassed("K") {
		if out == "C" {
			fmt.Printf("%gK er %g°C\n", fahr, conv.KelvinToCelsius(fahr))
		} else if out == "F" {
			fmt.Printf("%gK er %g°F\n", fahr, conv.KelvinToFahrenheit(fahr))
		} else if out == "K" {
			fmt.Printf("%gK\n", fahr)
		}
	}

	if isFlagPassed("funfacts") && isFlagPassed("t") {
		funfact := ""
		if strings.ToLower(funfactsArg) == "sun" {
			funfact = funfacts.SunFunFact(tempScale)
		} else if strings.ToLower(funfactsArg) == "luna" {
			funfact = funfacts.LunaFunFact(tempScale)
		} else if strings.ToLower(funfactsArg) == "terra" {
			funfact = funfacts.TerraFunFact(tempScale)
		} else {
			fmt.Println("Ugyldig funfacts valg, tillatte verdier: sun, luna, terra")
			return
		}

		fmt.Println(funfact)
	}
}

func isFlagPassed(s string) {
	panic("unimplemented")
}

		
