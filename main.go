package main

import (
	"Task27/pkg/storage"
	"Task27/pkg/student"
	"fmt"
)

func newStudent() {
	stydent := make(storage.StydentMap)
	for {
		studentTmp := student.Student{}
		fmt.Println("Введите через пробел имя студента, его возраст и курс.")
		n, err := fmt.Scan(&studentTmp.Name, &studentTmp.Age, &studentTmp.Grade)
		if n == 0 {
			stydent.Get()
			break
		} else if err != nil {
			fmt.Println("Ошибка:", err)
			break
		}
		stydent.Put(studentTmp)
	}
}

func main() {
	newStudent()
}
