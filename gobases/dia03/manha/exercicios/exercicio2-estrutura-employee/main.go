package main

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID           int
	Position     string
	PersonalData Person
}

func (e Employee) PrintEmployee() {
	fmt.Printf("Employee ID: %d\n", e.ID)
	fmt.Printf("Name: %s\n", e.PersonalData.Name)
	fmt.Printf("Date of Birth: %s\n", e.PersonalData.DateOfBirth)
	fmt.Printf("Position: %s\n", e.Position)
}

func main() {
	// Instanciando uma Person
	person := Person{
		ID:          1,
		Name:        "Eduardo Rumão",
		DateOfBirth: "1990-11-29",
	}

	// Instanciando um Employee, usando a pessoa acima
	employee := Employee{
		ID:           100,
		Position:     "CTO",
		PersonalData: person,
	}

	// Executando o método que imprime os dados do funcionário
	employee.PrintEmployee()
}
