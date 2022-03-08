/*
Задача № 2. Получить реверсную запись трехзначного числа
Пример:
вход: 346, выход: 643
вход: 100, выход: 001
*/
package main

import "fmt"

func main() {
	var number, first_digit, second_digit, third_digit, reverce_number int

	fmt.Println("Введите трехзначное число : ")
	fmt.Scanf("%d\n", &number)
	if 100 > number || number > 999 {
		fmt.Println("Введены неверные данные. Операция не может быть выполнена.")
	} else {
		third_digit = number % 10
		second_digit = number / 10 % 10
		first_digit = number / 100
		reverce_number = third_digit*100 + second_digit * 10 + first_digit
		fmt.Printf("Реверсная запись - %03d.\n", reverce_number)
	}
}
