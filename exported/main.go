package main

import (
	"log"
	"myapp/exported/staff"
)

//Exported vs Unexported
/*
create a variable up here at the package level called employees and it will be equal to a slice of staff dot employee.
And now I'll populate that with some data.
*/

var employees = []staff.Employee{

	/* I imported the type employee from the package staff. the reason I can do that is because employee this type
	begins with an uppercase letter and the only thing that's visible outside of this package are types, variables and functions
	that begin with an uppercase letter, that's really important to remember.
	*/
	{FirstName: "John", LastName: "Smith", Salary: 30000, FullTime: true},
	{FirstName: "Sally", LastName: "Jones", Salary: 50000, FullTime: true},
	{FirstName: "Mark", LastName: "Smithers", Salary: 60000, FullTime: true},
	{FirstName: "Alan", LastName: "Anderson", Salary: 15000, FullTime: false},
	{FirstName: "Margaret", LastName: "Carter", Salary: 100000, FullTime: false},
}

func main() {
	// initialise the package with some data.
	myStaff := staff.Office{
		AllStaff: employees,
	}

	//myStaff.All()

	// log.Println(myStaff.All())

	// Set new variable target make sure the varible is exported by starting with upper case leter.

	staff.OverPaidLimit = 70000

	log.Println("UPPER CASE exported", myStaff.Overpaid())
	log.Println("NOT UPPER case", myStaff.Underpaid())

}
