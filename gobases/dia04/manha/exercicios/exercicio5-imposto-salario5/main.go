package main

import (
	"errors"
	"fmt"
)

func calculateSalary(hoursWorked int, hourlyRate float64) (float64, error) {
	if hoursWorked < 80 {
		return 0, errors.New("error: the worker cannot have worked less than 80 hours per month")
	}
	if hoursWorked < 0 {
		return 0, errors.New("error: the worker cannot have worked less than 80 hours per month")
	}
	salary := float64(hoursWorked) * hourlyRate
	if salary >= 150000 {
		salary -= salary * 0.1 // Desconta 10%
	}
	return salary, nil
}

func main() {
	salary, err := calculateSalary(160, 1000) // variavel para teste
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Monthly salary: $%.2f\n", salary)
}
