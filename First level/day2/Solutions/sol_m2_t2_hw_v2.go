/*
Задача № 4. Проверить, является ли четырехзначное число палиндромом
Пример:
Вход: 1221  Выход: 1221 - палиндром
Вход: 1234  Выход: 1234 - не палиндром
*/

package main

import "fmt"

func main() {
	var number, first_digit, second_digit, third_digit, fourth_digit int
	fmt.Println("Введите четырехзначное число : ")
	fmt.Scanf("%d\n", &number)
	if 1000 > number || number > 9999 {
		fmt.Println("Введены неверные данные. Операция не может быть выполнена.")
	} else {
		fourth_digit = number % 10
		third_digit = number / 10 % 10
		second_digit = number / 100 % 10
		first_digit = number / 1000
		if fourth_digit == first_digit && third_digit == second_digit {
			fmt.Printf("Число %d - палиндром.\n", number)
		} else {

			fmt.Printf("Число %d - не палиндром.\n", number)
		}
	}
}
