package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type subject struct {
	name  string
	grade float64
}

func (sub subject) getSubjectName(r *bufio.Reader) string {
	fmt.Print("Enter the subjects name: ")

	subjectName, err := r.ReadString('\n')
	if err != nil {
		fmt.Print("subject name canot be empty try again: ")
		subjectName, err = r.ReadString('\n')
	} else {
		strings.Trim(subjectName, " ")
		if len(subjectName) == 0 {
			fmt.Print("subject name canot be empty try again: ")
			sub.getSubjectName(r)
		}
	}
	fmt.Println("this is the subjects name", subjectName)

	return subjectName
}

func (sub subject) getSubjectGrade() float64 {
	fmt.Print("Enter the grade out of 100 : ")
	var grade string
	fmt.Scan(&grade)
	grade_int, err := strconv.ParseFloat(grade, 64)
	for err != nil || grade_int < 0 || grade_int > 100 {
		fmt.Println(" ")
		fmt.Printf("The grade you entered ==> %v is invalid please enter a number from 0 to 100: ", grade)
		fmt.Scan(&grade)
		grade_int, err = strconv.ParseFloat(grade, 64)
	}

	return grade_int
}

type student struct {
	name     string
	subjects []subject
}

func (stdt student) getStudentName() string {
	fmt.Print("Enter Your name: ")
	var fname, lname string
	fmt.Scanln(&fname, &lname)
	for (len(fname) <= 2 || !IsLetterOnly(fname)) || (len(lname) <= 2 || !IsLetterOnly(lname)) {
		fmt.Println(" ")
		fmt.Printf("The name you entered ==> %s %s is invalid please enter a real name containing only letters: ", fname, lname)
		fmt.Scanln(&fname, &lname)

	}
	return fname + " " + lname
}

func (stdt student) displayStudentInfo() {
	fmt.Println("")
	fmt.Printf("%v's grade report\n", stdt.name)
	var sm float64
	fmt.Println("_________________________________________")
	fmt.Printf("|%-5v   |%-20v   |%-7v|\n", "no", "subject", "grade")
	fmt.Println("_________________________________________")

	for i, sub := range stdt.subjects {
		v := strings.Trim(sub.name, "\n")
		fmt.Printf("|%-5v   |%-20v   |%-7.3f|\n", i+1, v, sub.grade)
		fmt.Println("_________________________________________")
		sm += sub.grade
	}
	fmt.Printf("|%-15v|%-24.3f|\n", "total", sm)
	fmt.Println("_________________________________________")

	fmt.Printf("|%-15v|%-24.3f|\n", "average", sm/float64(len(stdt.subjects)))
	fmt.Println("_________________________________________")

}

func IsLetterOnly(s string) bool {
	for _, i := range s {
		if !unicode.IsLetter(i) {
			return false
		}
	}
	return true
}

func askChoice(r *bufio.Reader) {
	fmt.Println("")
	fmt.Println("Student Grade Calculator")
	fmt.Println("\tMenu")
	fmt.Println("\t 1,Calculate Grade\n\t 2,Exit ")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		std1 := student{}
		std1.name = std1.getStudentName()

		fmt.Print("Enter the number of subjects taken: ")
		var no_subjects string
		fmt.Scan(&no_subjects)
		number_of_subjects, err := strconv.Atoi(no_subjects)
		for err != nil || number_of_subjects <= 0 {
			fmt.Println(" ")
			fmt.Printf("The number of subjects you entered ==> %v is invalid please enter a number: ", no_subjects)
			fmt.Scan(&no_subjects)
			number_of_subjects, err = strconv.Atoi(no_subjects)

		}
		var sm float64
		for i := range number_of_subjects {
			fmt.Println("")
			sub := subject{}
			fmt.Println("subject ", i+1)
			sub.name = sub.getSubjectName(r)
			sub.grade = sub.getSubjectGrade()
			sm += sub.grade
			std1.subjects = append(std1.subjects, sub)

		}
		std1.displayStudentInfo()
		askChoice(r)
	case 2:
		fmt.Println("exiting...")
		break
	default:
		askChoice(r)

	}
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	askChoice(reader)

}
