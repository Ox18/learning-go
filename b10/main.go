package main

import "fmt"

func main() {
	// array
	var array [5]int

	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	array[4] = 5

	fmt.Println(array, len(array), cap(array))

	// Slice
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice, len(slice), cap(slice))

	// Métodos en el slice
	fmt.Println(slice[1:3])
	fmt.Println(slice[1:])
	fmt.Println(slice[:4])
	fmt.Println(slice[:])

	// Append
	slice = append(slice, 6)
	fmt.Println(slice)
}
