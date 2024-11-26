package main

import "fmt"

//Employee struct to hold employee details
type Employee struct {
	Name       string
	Department string
	Salary     float64
}

//UpdateSalary updates the salary of an employee
func UpdateSalary(emp *Employee, newSalary float64) {
	emp.Salary = newSalary

}
func main() {
	//Creating an instance of Employee
	emp := Employee{
		Name:       "John Doe",
		Department: "Software Engineer",
		Salary:     50000.0,
	}
	//Printing initial state
	fmt.Println("Before update:", emp)

	//Updating the salary
	fmt.Println(&emp)
	UpdateSalary(&emp, 60000.0)

	//Printing updated state
	fmt.Println("After update:", emp)

}
