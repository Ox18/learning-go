package main

func main() {
	modulo := 4 % 2

	switch modulo {
	case 0:
		println("Even")
	case 1:
		println("Odd")
	default:
		println("Unknown")
	}

	// other example
	switch modulo := 4 % 2; modulo {
	case 0:
		println("Even")
	case 1:
		println("Odd")
	default:
		println("Unknown")
	}

	// sin ninguna condición
	value := 200
	switch {
	case value > 100:
		println("Greater than 100")
	case value < 100:
		println("Less than 100")
	default:
		println("Equal to 100")
	}
}
