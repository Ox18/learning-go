package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Jose"] = 32
	m["Maria"] = 27

	fmt.Println(m)

	// recorrer map
	for k, v := range m {
		fmt.Println(k, v)
	}

	value, ok := m["Jose"]

	if ok {
		fmt.Println("Existe", value)
	} else {
		fmt.Println("No existe")
	}
}
