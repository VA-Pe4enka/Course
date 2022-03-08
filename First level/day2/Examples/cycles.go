package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Инициализация
	number := 7
	// Массив из 5 значений int
	var array [5]int 
	var str string = "hello"
	var character1 = 'h'
	var character2 rune = 'o'

	// Условный цикл
	for number > 0 {
		fmt.Print(number)
		// Изменение условий
		number--
	}
	fmt.Println()
	index := 0
	for index < len(array) {
		array[index] = index + 10
		index++
	}
	fmt.Println(array)

	for i := 0; i < len(array); i++ {
		fmt.Printf("%v ", array[i])
	}
	fmt.Println()
	// Возвращает индекс и значение элемента массива
	for index, value := range array {
		fmt.Printf("%v: %v\n", index, value)
	}
	// Используем только значение
	for  _, value := range array {
		fmt.Printf("%v\n", value)
	}
	
	array[4] = 5
	fmt.Println(array)

	fmt.Println(str, reflect.TypeOf(str))
	fmt.Println(character1, reflect.TypeOf(character1))
	for _, value := range str {
		fmt.Printf("%c ", value)
	}
	fmt.Println(len(str))
	s2 := "Привет"
	fmt.Println(len(s2))
	fmt.Printf("%v, %[1]c, %[1]U\n", character2)
	// s2[4] = 'a' // Error
	//Error, так нельзя
	// if "a" == 'a' {
	// 	fmt.Println("work")
	// }
	// Это работает
	if "a"[0] == 'a' {
		fmt.Println("work")
	}
}