/*
Задача № 2. Получить реверсную запись трехзначного числа
Пример: 
вход: 346, выход: 643
вход: 100, выход: 001
*/

package main

import "fmt"

func main() {
	var number, first, second, third int

    fmt.Println("Введите трехзначное число:")
	fmt.Scanf("%d\n", &number)

	first = number / 100
	second = (number / 10) % 10
	third = number % 10

	result := third*100 + second*10 + first

	fmt.Printf("%d", result)
}