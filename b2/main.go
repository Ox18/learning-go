package main

import "fmt"

func main() {
	// Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.14

	fmt.Println("El valor de pi es: ", pi)
	fmt.Println("El valor de pi2 es: ", pi2)

	// Declaración de variables enteras
	base := 12
	fmt.Println(base)

	var altura int = 14
	fmt.Println(altura)

	var area int

	fmt.Println(area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("El área del cuadrado es: ", areaCuadrado)
}
