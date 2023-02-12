package main

import "fmt"

func main() {
	// Defer
	defer fmt.Println("World")
	fmt.Println("Hello")
	// Cases para usar defer:
	// - Cerrar el archivo
	// - Cerrar la conexión a la base de datos

	// Continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		// Continue
		if i == 5 {
			fmt.Println("es 5")
			continue
		}

		// Break
		if i == 8 {
			fmt.Println("es 8")
			break
		}
	}
}
