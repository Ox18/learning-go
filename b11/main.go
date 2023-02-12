package main

import "fmt"

func isPalindromo(a string) bool {
	for i := 0; i < len(a)/2; i++ {
		if a[i] != a[len(a)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	slice := []string{"hola", "que", "tal"}

	for i := range slice {
		fmt.Println(slice[i])
	}

	if isPalindromo("hola") {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}

	if isPalindromo("oso") {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}
