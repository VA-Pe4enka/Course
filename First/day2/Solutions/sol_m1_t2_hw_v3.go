/*
Задача № 2. Получить реверсную запись трехзначного числа
Пример:
вход: 346, выход: 643
вход: 100, выход: 001
*/
package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	//Вариант с проверкой на трёхзначность, с индексацией строки:
	var number string

	fmt.Print("Three-digit number:\n>>> ")
	_, err := fmt.Scan(&number)
	if err != nil {
		log.Fatal("wrong data")
	}

	_, err = strconv.Atoi(number) // strconv.ParseInt()
	fmt.Printf("%v, %T\n", number, number)
	if len(number) != 3 || err != nil {
		log.Fatal("not a three-digit number")
	}

	fmt.Print(string(number[2]) + string(number[1]) + string(number[0]))
	fmt.Printf("%c\n", number[2]+number[1]+number[0])

	// Вариант без проверки на трёхзначность:
	var x, y, z int

	fmt.Print("Three-digit number:\n>>> ")
	_, err = fmt.Scanf("%1d%1d%1d", &x, &y, &z)
	if err != nil {
		log.Fatalln("Not a number")
	}

	fmt.Printf("%1d%1d%1d", z, y, x)
}
