# GOLANG-LEARNING
Install Git-bash as terminal for Windows
Install go from https://go.dev/dl/
Install Visual Studio for any platform
After Visual Studio installed need to intall Go modules
	Crtl + Shift + P 
	> Go: Install/update (select all)

Create Project folder

Enter project folder via Git-Bash command line
Get your GOPATH by following command 
	$ echo $GOPTAH
Initiate Go modules in project folder
	$ go mod init 'myapp'  'name of your application or location like github.com/user/myapp'

Opening Settings short cup
	Crtl + ,
	File --> Preferences --> Settings
	Update Settings
		"terminal.integrated.defaultProfile.windows": "Git Bash",
		"terminal.external.windowsExec": "C:\\Program Files\\Git\\git-bash.exe",
		"terminal.integrated.cwd": "C:\\Users\\ricar\\Desktop\\Nextcloud\\Programming\\GitServer\\golang",

Download godoc modules for references
	$ go get golang.org/x/tools/cmd/godoc
	$ go install golang.org/x/tools/cmd/godoc
To view godoc initial webserver
	$ godoc --http=:6060
Open favour browser browser and open URL
	http://localhost:6060
Very first line in all GO files start with package
	package main
Every main package has to have a function called main
	func main() {

	}
Example of Hello world in main function
	package main

	import "fmt"

	func main() {
		fmt.Println("Hello World!")
	}

Run go program through Git-Bash command
	$ go run helloWorld/main.go 

BUILD a go executable program through Git-Bash command
	$ go build -o hello.exe helloWorld/main.go

Variables known as short and operators

	## Standard string varible....
	var sayWhat string
	sayWhat = "Hello World, through function & variable"

	## Short hand string varible.
	sayWhat := "Hello World, through function & short hand variable !!!"
	sayHelloWorld(sayWhat)

	## private vs Public
	var privateVar = "I am private"                // started with lowecase to be private
	var PublicVar = "I am public (or exported)"    // started with UPPER case to be public

string array in GO is called [] slice to a slice of string
2d array in GO is called MAP to store key & values

Const is constant varible outside of the function and can never change or never has to be used

Scope refer to wher ein your program that you have declared a varible

Accessing Public vs private functions.

	func notExported(){
		//Start with lower case can not be accessed
	}

	func Exported() {
		// Starts with Upper case so ti CAN be accessed
	}

Lesson 20 using 3rd party package

Ignore the second inut "error" fromt he function with " _ " 
	userInput, _ := reader.ReadString('\n')
Delete new line from string for windows -1 = everywhere
	userInput = strings.Replace(userInput, "\r\n", "", -1)
Delete new line from string for unix/linux/darwin -1 = everywhere
	userInput = strings.Replace(userInput, "\n", "", -1)   //


MAP is a data structure that has a index & a value (String Array)
	coffees := make(map[int]string)
	coffees[1] = "Cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"

Saved in string-interpolation find file fmt-printing-cheatsheet.pdf

String-challenge code has example of bool YES or NO

## Section 4: Types, Expressions and Composition

## TYPES
Basic types - numbers, strings, booleans
Aggredate types - arrays, structs
Reference types - Pointers, slices, maps, fucntions, channels
 -> SLICES are arrays with additional functionalities 
Interface type 

## EXPRESSIONS
Boolean if 
Compound boolean if 
Switch()

## COMPOSITION
Composition vs Inheritance
Go is not OOP 
Example: car in OOP, car with Composition



######################
# AREK STOP AT 25
# RICO START 34 MAPS
######################
