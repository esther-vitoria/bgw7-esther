package main

import (
	"errors"
	"fmt"
)

type SalaryError struct{}

func (e *SalaryError) Error() string {
	return "Error: salary is less than 10000"
}

var ErrSalary = &SalaryError{}

func main() {
	salary := 9000 // variavel para teste
	var err error
	if salary <= 10000 {
		err = ErrSalary
	}

	if errors.Is(err, ErrSalary) {
		fmt.Println(err)
		return
	}

	fmt.Println("Salary ok")
}
