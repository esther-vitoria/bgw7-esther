package main

import (
	"errors"
	"fmt"
)

var ErrSalary = errors.New("error: salary is less than 10000")

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
