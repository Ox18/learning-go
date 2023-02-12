package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	worldMessage := "Hola, Mundo!"
	helloMessage := "Hello, World!"

	fmt.Println(worldMessage, helloMessage)
	fmt.Println("Hello, World!")

	// PrintF
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)

	// SprintF
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	fmt.Println(message)

	// Tipos de datos
	var entero8 int8 = 127
	fmt.Printf("Entero8: %T", entero8)
}
