package task4

import (
	"crypto/sha256"
	"fmt"
)

type Student struct {
	rollnumber int
	name       string
	address    string
	subject    []string
}

func NewStudent(roll int, name string, address string, subs []string) *Student {
	s := new(Student)
	s.rollnumber = roll
	s.name = name
	s.address = address
	s.subject = subs
	return s
}

type StudentList struct {
	list []*Student
}

func (ls *StudentList) CreateStudent(roll int, name string, address string, subs []string) *Student {
	st := NewStudent(roll, name, address, subs)
	ls.list = append(ls.list, st)
	return st
}

func CalculateHash(st []*Student) string {
	var stringToHash string
	for _, val := range st {
		for _, sub := range val.subject {
			stringToHash += sub
		}
	}
	fmt.Println("string received : \n", stringToHash)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(stringToHash)))
}

func task4() {
	student := new(StudentList)
	subs1 := []string{"BlockChain", "Golang"}
	subs2 := []string{"MMDS", "InfoSec"}
	student.CreateStudent(24, "Faizan", "Joiya", subs1)
	student.CreateStudent(25, "Naveed", "Rawalpindi", subs2)
	output := CalculateHash(student.list)
	fmt.Println("Hash :", output)

}
