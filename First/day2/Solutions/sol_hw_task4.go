/*
Задача № 4. Проверить, является ли четырехзначное число палиндромом
Пример:
Вход: 1221  Выход: 1221 - палиндром
Вход: 1234  Выход: 1234 - не палиндром
*/

package main

import "fmt"

func main() {
	var number, num1, num2, num3, num4 int

	fmt.Println("Введите четырехзначное число:")
	fmt.Scanf("%d\n", &number)

	num1 := number / 1000
	num2 := (number / 100) % 10
	num3 := (number / 10) % 10
	num4 := number % 10

	if num1 == num4 && num2 == num3 {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Не палиндром")
	}
}
