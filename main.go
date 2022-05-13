package main

import (
	"Task27/Struct"
	"fmt"
)

type stydentMap map[string]*Struct.Student

func (s stydentMap) Put(a Struct.Student) {
	s[a.Name] = &a
}

func (s stydentMap) Get() {
	fmt.Println("Данные студентов:")

	for _, value := range s {
		fmt.Printf("%-9s %-3v %v\n", value.Name, value.Age, value.Grade)
	}
}

func newStudent() {
	stydent := make(stydentMap)
	for {
		a := Struct.Student{}
		fmt.Println("Введите через пробел имя студента, его возраст и курс.")
		n, err := fmt.Scan(&a.Name, &a.Age, &a.Grade)
		if n == 0 {
			stydent.Get()
			//printStudent()
			break
		} else if err != nil {
			fmt.Println("Ошибка:", err)
			break
		}
		stydent.Put(a)
	}
}

func main() {
	newStudent()
}
