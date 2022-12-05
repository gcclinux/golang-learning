package staff

// So when it begins with an uppercase letter, it's exported and visible outside of the package.
//
// But when it begins with a lowercase letter, it's not exported and is not visible

/*
	Again, types, variables, constants and functions that begin with an UPPER case letter are exported,
	so package level variables with an UPPER case letter are also exported.
	Those that begin with a lowercase letter are not exported.
*/

var OverPaidLimit = 75000
var underpaidMinimum = 20000

type Employee struct {
	FirstName string
	LastName  string
	Salary    int
	FullTime  bool
}

type Office struct {
	// Slice for all Staff that includes employees
	AllStaff []Employee
}

func (e *Office) All() []Employee {
	return e.AllStaff
}

func (e *Office) Overpaid() []Employee {
	var overpaid []Employee

	for _, x := range e.AllStaff {
		if x.Salary > OverPaidLimit {
			overpaid = append(overpaid, x)
		}
	}
	return overpaid
}

func (e *Office) Underpaid() []Employee {
	var underpaid []Employee

	for _, x := range e.AllStaff {
		if x.Salary < underpaidMinimum {
			underpaid = append(underpaid, x)
		}
	}
	return underpaid
}
