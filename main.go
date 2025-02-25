package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Student struct {
	FirstName       string
	LastName        string
	MathScore       float32
	ScienceScore    float32
	LiteratureScore float32
	Average         float32
}

var students []Student

func main() {
	showMenu()
}

func showMenu() {
	var choiceOfUser string
	fmt.Println("\nStudent Management System Menu:")
	fmt.Println("1. Add a student")
	fmt.Println("2. Display students")
	fmt.Println("3. Identify and display the top-performing student")
	fmt.Println("4. Sort Students by their average scores")
	fmt.Println("5. Save student data in a text file")
	fmt.Println("6. Exit")
	fmt.Print("Please enter your choice: ")
	fmt.Scanln(&choiceOfUser)

	switch choiceOfUser {
	case "1":
		addStudent()
	case "2":
		displayStudents()
	case "3":
		displayTopPerformingStudent()
	case "4":
		sortStudents()
	case "5":
		saveStudentsDataInJsonFile()
	case "6":
		return
	}

}

func addStudent() {
	var firstName, lastName string
	var mathScore, scienceScore, literatureScore float32
	fmt.Print("\nEnter the First Name: ")
	fmt.Scanln(&firstName)
	fmt.Print("Enter the Last Name: ")
	fmt.Scanln(&lastName)
	fmt.Print("Enter the Math Score: ")
	fmt.Scanln(&mathScore)
	fmt.Print("Enter the Science Score: ")
	fmt.Scanln(&scienceScore)
	fmt.Print("Enter the Literature Score: ")
	fmt.Scanln(&literatureScore)

	newStudent := Student{
		FirstName:       firstName,
		LastName:        lastName,
		MathScore:       mathScore,
		ScienceScore:    scienceScore,
		LiteratureScore: literatureScore,
	}

	students = append(students, newStudent)

	go calcAverage(&students[len(students)-1])

	showMenu()
}

func displayStudents() {
	for i, s := range students {
		fmt.Printf("%d. %s %s | Math: %.2f | Science: %.2f | Literature: %.2f | Average: %.2f \n", i+1, s.FirstName,
			s.LastName, s.MathScore, s.ScienceScore, s.LiteratureScore, s.Average)
	}

	showMenu()
}

func calcAverage(s *Student) {
	s.Average = (s.MathScore + s.ScienceScore + s.LiteratureScore) / 3
}

func sortStudents() {
	sort.SliceStable(students, func(i, j int) bool {
		return students[i].Average > students[j].Average
	})

	fmt.Println("\nStudents have been sorted by their average scores")
	showMenu()
}

func displayTopPerformingStudent() {
	if len(students) == 0 {
		fmt.Println("\nYou have not saved any students.")
	} else {
		sort.SliceStable(students, func(i, j int) bool {
			return students[i].Average > students[j].Average
		})

		s := students[0]
		fmt.Printf("\n%s %s | Math: %.2f | Science: %.2f | Literature: %.2f | Average: %.2f \n", s.FirstName,
			s.LastName, s.MathScore, s.ScienceScore, s.LiteratureScore, s.Average)
	}
	showMenu()
}

func saveStudentsDataInJsonFile() {

	if len(students) == 0 {
		fmt.Println("\nYou have not saved any students.")
	} else {
		jsonData, err := json.Marshal(students)
		if err != nil {
			panic(err)
		}
		err = os.WriteFile("students.json", jsonData, os.ModePerm)
		if err != nil {
			panic(err)
		}

		fmt.Println("\nStudents have been saved in a text file")
	}
	showMenu()
}
