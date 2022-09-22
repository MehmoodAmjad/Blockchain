package main

import "fmt"

type Student struct {
	rollnumber int
	name       string
	address    string
}

func NewStudent(roll int, name string, address string) *Student {
	s := new(Student)
	s.rollnumber = roll
	s.name = name
	s.address = address
	return s
}

type StudentList struct {
	list []*Student
}

func (ls *StudentList) CreateStudent(roll int, name string, address string) *Student {
	st := NewStudent(roll, name, address)
	ls.list = append(ls.list, st)
	return st
}
func main() {
	student := new(StudentList)
	student.CreateStudent(24, "Faizan", "Joiya")
	student.CreateStudent(25, "Naveed", "Rawalpindi")
	fmt.Println(student)
}
