package main

import (
	"fmt"
	hello "unit-testing/hello"
)

func main() {

	p := hello.NewPerson("")

	s, err := p.Salute()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(s)
}
