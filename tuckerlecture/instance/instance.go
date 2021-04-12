package main

import "fmt"

type Student struct {
	name  string
	age   int
	grade int
}

func (t *Student) SetName(name string) {
	t.name = name
}

func (t *Student) SetAge(age int) {
	t.age = age
}

func (t *Student) PrintStudentInfo() {
	fmt.Println(t.name, t.age, t.grade)
}

func SetName(t *Student, name string) {
	t.name = name
}

func PrintStudentInfo(t *Student) {
	fmt.Println(t.name, t.age, t.grade)
}

func main() {
	var a *Student
	a = &Student{name: "junhaeng", age: 32, grade: 100}
	a.PrintStudentInfo()

	// OOP 형태(Subject, Verb, Object)로 정의
	// 컴퓨터 입장에서는 Procedure나 OOP나 동일하지만
	// 프로그래머에게는 완전히 새로운 패러다임(!)
	// Entity 간의 Relationship이 중요해졌다

	// 프로시저 기반
	// 컴퓨터 입장에서는 Procedure든 OOP든 차이가 없다(!)
	// 아래 a.SetName과 완전히 동일함
	SetName(a, "jihee")
	PrintStudentInfo(a)

	// 객체 기반
	a.SetName("helloworld") // Subject역할을 하는 a를 Instance라고 하자(일종의 생명주기)
	a.PrintStudentInfo()
}
