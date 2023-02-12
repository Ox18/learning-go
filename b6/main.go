package main

import "fmt"

func main() {
	// For conditional
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// For forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
		if counterForever == 10 {
			break
		}
	}

	listaNumerosPares := []int{2, 4, 6, 8, 10}

	// For range
	for index, value := range listaNumerosPares {
		fmt.Println(index, value)
	}
}
