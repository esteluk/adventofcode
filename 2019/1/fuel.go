package main

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Fuel Calculates the fuel required for a given mass of module
func Fuel(mass int) int {
	return (mass / 3) - 2
}

// FuelForFuel makes sure that our fuel has enough fuel
func FuelForFuel(mass int) int {
	var fuel int = 0

	var f int = Fuel(mass)
	for f > 0 {
		fuel += f
		f = Fuel(f)
	}
	fuel += Max(0, f)
	return fuel
}
