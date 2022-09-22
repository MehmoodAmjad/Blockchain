package main

import (
	"fmt"
	"strings"
)

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
func Display(st []*Student) {
	for index, val := range st {
		fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), index, strings.Repeat("=", 25))
		fmt.Println("Student Roll No", val.rollnumber)
		fmt.Println("Student Name", val.name)
		fmt.Println("Student Address", val.address)
	}
}
func main() {
	student := new(StudentList)
	student.CreateStudent(24, "Faizan", "Joiya")
	student.CreateStudent(25, "Naveed", "Rawalpindi")
	Display(student.list)
}
