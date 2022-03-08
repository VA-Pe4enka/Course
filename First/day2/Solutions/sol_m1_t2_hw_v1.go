/*
Задача № 2. Получить реверсную запись трехзначного числа
Пример:
вход: 346, выход: 643
вход: 100, выход: 001
*/
package main

import "fmt"

func main() {
	var number int
	fmt.Println("Введите трехзначное число")
	fmt.Scanf("%d", &number)
	first := number / 100
	second := number / 10 % 10
	last := number % 10
	fmt.Printf("%d%d%d\n", last, second, first)

}
