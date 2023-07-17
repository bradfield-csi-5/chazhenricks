package main

import "fmt"

type Employee struct {
	ID       int
	Name     string
	Position string
}

func main() {
	dilbert := Employee{ID: 1, Name: "dilbert", Position: "Developer"}
	chaz := Employee{ID: 2, Name: "Chaz", Position: "Developer"}

	var employeeOfMonth *Employee = &dilbert
	otherEmployee := chaz

	fmt.Printf("%s\n", employeeOfMonth.Name)
	fmt.Printf("%s\n", otherEmployee.Name)
}
