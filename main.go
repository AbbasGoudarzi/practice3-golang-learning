package main

import (
	"fmt"
)

type student struct {
	firstName       string
	lastName        string
	mathScore       float32
	scienceScore    float32
	literatureScore float32
	average         float32
}

var students []student

func main() {
	showMenu()
}

func showMenu() {
	var choiceOfUser string
	fmt.Println("\nStudent Management System Menu:")
	fmt.Println("1. Add a student")
	fmt.Println("2. Display students")
	fmt.Println("3. Exit")
	fmt.Print("Please enter your choice: ")
	fmt.Scanln(&choiceOfUser)

	switch choiceOfUser {
	case "1":
		addStudent()
	case "2":
		displayStudents()
	case "3":
		return
	}

}

func addStudent() {
	var firstName, lastName string
	var mathScore, scienceScore, literatureScore float32
	fmt.Print("Enter the First Name: ")
	fmt.Scanln(&firstName)
	fmt.Print("Enter the Last Name: ")
	fmt.Scanln(&lastName)
	fmt.Print("Enter the Math Score: ")
	fmt.Scanln(&mathScore)
	fmt.Print("Enter the Science Score: ")
	fmt.Scanln(&scienceScore)
	fmt.Print("Enter the Literature Score: ")
	fmt.Scanln(&literatureScore)

	newStudent := student{
		firstName:       firstName,
		lastName:        lastName,
		mathScore:       mathScore,
		scienceScore:    scienceScore,
		literatureScore: literatureScore,
	}

	students = append(students, newStudent)

	go calcAverage(&students[len(students)-1])

	showMenu()
}

func displayStudents() {
	for i, s := range students {
		fmt.Printf("%d. %s %s | Math: %.2f | Science: %.2f | Literature: %.2f | Average: %.2f \n", i+1, s.firstName,
			s.lastName, s.mathScore, s.scienceScore, s.literatureScore, s.average)
	}

	showMenu()
}

func calcAverage(s *student) {
	s.average = (s.mathScore + s.scienceScore + s.literatureScore) / 3
}
