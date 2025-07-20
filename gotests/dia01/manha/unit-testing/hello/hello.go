package hello

import (
	"fmt"
)

type Person struct {
	Name string
}

func NewPerson(name string) Person {
	return Person{
		Name: name,
	}
}

func (p *Person) Salute() (string, error) {

	if p.Name == "" {
		return "", fmt.Errorf("error, name is empty")
	}
	msg := fmt.Sprintf("Hi, %s.\nWelcome back to Bootcamp GO!!", p.Name)
	return msg, nil
}
