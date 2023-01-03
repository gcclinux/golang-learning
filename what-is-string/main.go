package main

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	myString := "This is a clear example of why we search in one case only"

	searchString := strings.ToLower(myString)

	if strings.Contains(searchString, "this") {
		fmt.Println("Found it")
	} else {
		fmt.Println("Did not find it")
	}

	// other case functions
	fmt.Println(strings.ToLower(myString))
	fmt.Println(strings.ToUpper(myString))

	//deprecated method
	//fmt.Println(strings.Title(strings.ToLower(myString)))

	//alternative method
	fmt.Println(cases.Title(language.Und).String(myString))
}
