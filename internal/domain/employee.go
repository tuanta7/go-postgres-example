package domain

import "time"

type EmployeeGender string

const (
	EmployeeGenderMale   EmployeeGender = "M"
	EmployeeGenderFemale EmployeeGender = "F"
)

// Employee represents an employee independently of database access types.
// BirthDate and HireDate represent calendar dates without a time of day.
type Employee struct {
	ID        int64
	BirthDate time.Time
	FirstName string
	LastName  string
	Gender    EmployeeGender
	HireDate  time.Time
}
