package main

import (
	"Task27/Struct"
	"fmt"
)

var studentMap = make(map[string]*Struct.Student, 10)

func newStudent() {
	for {
		a := Struct.Student{}
		fmt.Println("Введите через пробел имя студента, его возраст и курс.")
		n, err := fmt.Scan(&a.Name, &a.Age, &a.Grade)
		if n == 0 {
			printStudent()
			break
		} else if err != nil {
			fmt.Println("Ошибка:", err)
			break
		}
		studentMap[a.Put()] = &a
	}
}

func main() {
	newStudent()
}

func printStudent() {
	fmt.Println("Данные студентов:")
	for _, value := range studentMap {
		fmt.Printf("%-9s %-3v %v\n", value.Name, value.Age, value.Grade)
	}
}
