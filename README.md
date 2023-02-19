package main

import (
	"math"
	"reflect"
	"testing"
)

// Konverterer Fahrenheit til Celsius
func FahrenheitToCelsius(value float64) float64 {
	return (value - 32) * 5 / 9
}

// Konverterer Celsius til Fahrenheit
func CelsiusToFahrenheit(value float64) float64 {
	return value*9/5 + 32
}

// Konverterer Kelvin til Fahrenheit
func KelvinToFahrenheit(value float64) float64 {
	return value*9/5 - 459.67
}

// Konverterer Fahrenheit til Kelvin
func FahrenheitToKelvin(value float64) float64 {
	return (value + 459.67) * 5 / 9
}

// Konverterer Celsius til Kelvin
func CelsiusToKelvin(value float64) float64 {
	return value + 273.15
}

// Konverterer Kelvin til Celsius
func KelvinToCelsius(value float64) float64 {
	return value - 273.15
}

func withinTolerance(a, b, error float64) bool {
	// Først sjekk om tallene er nøyaktig like
	if a == b {
		return true
	}

	difference := math.Abs(a - b)

	// Siden vi skal dele med b, må vi sjekke om den er 0
	// Hvis b er 0, returner avgjørelsen om d er mindre enn feilmarginen
	// som vi aksepterer
	if b == 0 {
		return difference < error
	}

	// Tilslutt sjekk den relative differanse mot feilmargin
	return (difference/math.Abs(b)) < error
}

func TestFahrenheitToCelsius(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{32, 0},
		{50, 10},
		{212, 100},
	}

	for _, tc := range tests {
		got := FahrenheitToCelsius(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
		}
	}
}

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{0, 32},
		{10, 50},
		{100, 212},
	}

	for _, tc := range tests {
		got := CelsiusToFahrenheit(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
		}
	}
}

func TestKelvinToFahrenheit(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{273.15, 32},
		{283.15, 50},
		{373.15, 212},
	}

	for _, tc := range tests {
		got := KelvinToFahrenheit(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
		}
	}
}

func TestFahrenheitToKelvin(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{32, 273.15},
		{-40, 233.15},
		{212, 373.15},
	}

	for _, tc := range tests {
		got := FahrenheitToKelvin(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
		}
	}
}

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{0, 32},
		{-40, -40},
		{100, 212},
	}

	for _, tc := range tests {
		got := CelsiusToFahrenheit(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
		}
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{0, 273.15},
		{-273.15, 0},
		{100, 373.15},
	}

	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
		}
	}
}

func TestKelvinToFahrenheit(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{0, -459.67},
		{273.15, 32},
		{373.15, 212},
	}

	for _, tc := range tests {
		got := KelvinToFahrenheit(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
		}
	}
}

func TestKelvinToCelsius(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{0, -273.15},
		{273.15, 0},
		{373.15, 100},
	}

	for _, tc := range tests {
		got := KelvinToCelsius(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
		}
	}
}

