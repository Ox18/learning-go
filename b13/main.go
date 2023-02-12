package main

type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Ford", year: 2015}

	// myCar is a value of type car
	println(myCar.brand, myCar.year)

	// other form to instance class
	var otherCar car
	otherCar.brand = "Fiat"
	otherCar.year = 2010

	println(otherCar.brand, otherCar.year)
}
