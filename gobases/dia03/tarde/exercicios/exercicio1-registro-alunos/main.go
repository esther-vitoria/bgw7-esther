package main

import "fmt"

type Students struct {
	FirstName     string
	LastName      string
	ID            int
	AdmissionDate string
}

// Esse método substituí os dados fornecidos pelos alunos
func (s *Students) UpdateStudent(newStudent Students) {
	s.FirstName = newStudent.FirstName
	s.LastName = newStudent.LastName
	s.ID = newStudent.ID
	s.AdmissionDate = newStudent.AdmissionDate
}

func (s Students) Information() {
	fmt.Printf("Student\n-Name: %s\n-Last Name: %s\n-ID: %d\n-Admission date: %s\n", s.FirstName, s.LastName, s.ID, s.AdmissionDate)
}

func main() {
	student := Students{
		FirstName:     "José",
		LastName:      "Ulisses",
		ID:            1,
		AdmissionDate: "20/05/2025",
	}

	student.Information()

	newStudent := Students{
		FirstName:     "Esther",
		LastName:      "Balbino",
		ID:            2,
		AdmissionDate: "20/05/2025",
	}

	// Atua;liza o estudante para o novo estudante!
	student.UpdateStudent(newStudent)

	student.Information()
}
