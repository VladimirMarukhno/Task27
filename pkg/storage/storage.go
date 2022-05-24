package storage

import (
	"Task27/pkg/student"
	"fmt"
)

type StydentMap map[string]*student.Student

func (s StydentMap) Put(a student.Student) {
	s[a.Name] = &a
}

func (s StydentMap) Get() {
	fmt.Println()
	fmt.Println("Данные студентов:")
	for _, value := range s {
		fmt.Printf("%-9s %-3v %v\n", value.Name, value.Age, value.Grade)
	}
}
