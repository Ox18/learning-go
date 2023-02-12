package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, World!")

	x := 10
	y := 50

	// Sum
	result := x + y
	fmt.Println(result)

	// Rest
	result = y - x
	fmt.Println(result)

	// Multiplication
	result = x * y
	fmt.Println(result)

	// Division
	result = y / x
	fmt.Println(result)

	// Modulo
	result = y % x
	fmt.Println(result)

	// Increment
	x++
	fmt.Println(x)

	// Decrement
	y--
	fmt.Println(y)

	// Area de un rectangulo
	base := 10
	height := 5

	areaRectangulo := base * height

	fmt.Println(areaRectangulo)

	// Area de un trapecio
	baseMayor := 10
	baseMenor := 5
	heightTrapecio := 7

	areaTrapecio := (baseMayor + baseMenor) * heightTrapecio / 2

	fmt.Println(areaTrapecio)

	// Area de un circulo
	radio := 5
	areaCirculo := math.Pi * math.Pow(float64(radio), 2)

	fmt.Println(areaCirculo)
}
