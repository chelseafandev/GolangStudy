package main

import "fmt"

type Student struct {
	name  string
	age   int
	grade string
	class string
}

func (s *Student) PrintSungjuk() {
	fmt.Println(s.class, s.grade)
}

func (s *Student) InputSungjuk(class string, grade string) {
	s.class = class
	s.grade = grade
}

func main() {
	var student Student = Student{name: "Junhaeng", age: 32, class: "Math", grade: "B"}
	student.PrintSungjuk()

	student.InputSungjuk("Math", "A+")
	student.PrintSungjuk()
}
