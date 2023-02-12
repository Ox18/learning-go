package main

import (
	"fmt"

	"./mypackage"
)

func main() {
	var myCar mypackage.CarPublic

	myCar.Brand = "Ford"
	myCar.Year = 2018

	println(myCar.Brand)

	fmt.Println(myCar)
}
