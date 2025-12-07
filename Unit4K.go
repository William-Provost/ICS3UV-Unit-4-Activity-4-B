// Author: Mr Coxall
// Version: 1.0.0
// Date: 2025-12-07
// Fileoverview: This program shows the random built in function.

package main

import (
	"fmt"
	"math/rand/v2" // import the random package, new in Go Go 1.22+
)

func main() {
	// constants and variables
	var someRandomFloat float64 = 0.0
	var someRandomInteger int = 0
	var someRandomIntegerBetween10And20 int = 0

	// get a random float between 0 and 1
	someRandomFloat = rand.Float64()

	// get a random integer between 0 and 100
	// rand.Float64() generates [0, 1). Multiplying by 101 gives [0, 101). Converting to int truncates the decimal.
	someRandomInteger = int(rand.Float64() * 101)

	// get a random integer between 10 and 20
	// rand.Float64() * 11 gives [0, 11). Converting to int gives [0, 10]. Adding 10 gives [10, 20].
	someRandomIntegerBetween10And20 = int(rand.Float64() * 11) + 10

	// display results
	fmt.Println("Here are some random numbers:")
	fmt.Println(someRandomFloat)
	fmt.Println(someRandomInteger)
	fmt.Println(someRandomIntegerBetween10And20)

	fmt.Println("\nDone.")
}
