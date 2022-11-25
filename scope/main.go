package main

import (
	"myapp/scope/packageone"
)

/* package level variable lesson 16
var one = "One"
*/

var myVar = "Package Level variable"

func main() {
	/* Lesson 16
	var one = "this is blok level varibale"
	fmt.Println(one)
	myFunc()
	newString := packageone.PublicVar
	fmt.Println("From packageone:", newString)
	packageone.Exported()
	*/

	// Lesson 17
	// Declare package level variables for the main package myVar
	// Declare block level varibales for the main func blockVar
	// Declare package level variable in packageone named PackageVar
	// Craete and exported functions in packageone called PrintMe
	// Printout value in single line for myVar, blockVar, PackageVar

	var blockVar = "Block Level variable"
	packageone.PrintMe(myVar, blockVar)
}

/* Lesson 16
func myFunc() {
	fmt.Println(one)
}
*/
