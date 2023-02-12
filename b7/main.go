package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	fmt.Println("Hello, world.")

	valor := 1
	valor2 := 2

	if valor == 1 {
		fmt.Println("Valor é igual a 1")
	} else if valor == 2 {
		fmt.Println("Valor é igual a 2")
	} else {
		fmt.Println("Valor é diferente de 1 e 2")
	}

	if valor == 1 && valor2 == 2 {
		fmt.Println("Valor é igual a 1 e valor2 é igual a 2")
	}

	// Convertir textos a numeros
	value, err := strconv.Atoi("x")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(value)
}
