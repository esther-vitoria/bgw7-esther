package main

import (
	"fmt"
)

func main() {
	salary := 149000 // variavel para teste

	var err error
	if salary < 150000 {
		err = fmt.Errorf("error: the minimum taxable amount is 150,000 and the salary entered is: %d", salary)
		fmt.Println(err)
		return
	}

	fmt.Println("Must pay tax")
}
