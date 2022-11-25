package packageone

import (
	"fmt"
)

/* Lesson 16

var privateVar = "I am private"             // started with lowecase to be private
var PublicVar = "I am public (or exported)" // started with UPPER case to be public

func notExported() {
	// this is only available inside package.go
}

func Exported() {
	// this is available everywhere within the program
}

*/

// Lesson 17
var PackageVar = "Package Variable String in packageone"

func PrintMe(myVar, blockVar string) {
	fmt.Println("Printing: myVar:", myVar, ",blockVar:", blockVar, ",PackageVar:", PackageVar)
}
