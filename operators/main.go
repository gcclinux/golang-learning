package main

import (
	"fmt"
	"math"
)

func main() {
	// Using parentices can determine which part is calculated first

	answer := 7 + 3*4
	fmt.Println("Answer I:", answer)

	answer = (7 + 3) * 4
	fmt.Println("Answer II:", answer)

	fmt.Println()
	var radius = 12.0
	area := math.Pi * radius * radius
	fmt.Println("Area is: ", area)

	fmt.Println()
	// Integer division
	half := 1 / 2
	fmt.Println("Half with int division:", half)

	halffloat := 1.0 / 2.0
	fmt.Println("Halffloat with int division:", halffloat)

	fmt.Println()
	// square (raise to the power)
	badThreeSquared := 3 ^ 2
	fmt.Println("badThreeSquared:", badThreeSquared)

	goodThreeSquared := math.Pow(3.0, 2.0)
	fmt.Println("goodThreeSquared:", goodThreeSquared)
	fmt.Println()
	// modulues
	remainder := 50 % 3
	fmt.Println("Reminder:", remainder)

	// unary operators
	x := 3
	x++
	fmt.Println("X is now,", x)
	x--
	x--
	fmt.Println("X is now,", x)

	// Precedence
	// Multiplications in precedence
	fmt.Println()
	a := 12.0 * 3.0 / 4.0
	b := (12.0 * 3.0) / 4.0
	c := 12.0 * (3.0 / 4.0)
	fmt.Println("Precedence a:", a, "b:", b, "c:", c)

	// Integer Division NOT float
	unclear := 12 * (3 / 4)
	fmt.Println("Unclear:", unclear)

	// Float Division NOT Integer
	fmt.Println()
	f := 12.0 / 3.0 / 3.0
	fmt.Println("F:", f)

	// Parenthesis always takes priority / alwasy calculated first
	f = 12.0 / (3.0 / 4.0)
	fmt.Println("F:", f)

	// Addition & substraction
	fmt.Println()
	w := 12 + 3 - 4
	y := (12 + 3) - 4
	z := 12 + (3 - 4)
	fmt.Println("Precedence w:", w, "y:", y, "z:", z)

	w = 12 + 3*4
	y = (12 + 3) * 4
	z = 12 + (3 * 4)

	fmt.Println("Precedence w:", w, "y:", y, "z:", z)

	//does one number divide exactly into another?
	fmt.Println()
	xx := 12
	yy := 3
	if xx%yy == 0 {
		fmt.Println(y, "divided exactly into, x")
	} else {
		fmt.Println(y, "Does not divided exactly into, x")
	}
	fmt.Println()

	for m := 1; m <= 12; m++ {
		fmt.Println("The month after", m, "is", m%12+1)
	}

}
