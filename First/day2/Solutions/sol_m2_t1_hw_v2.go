/*
Задача № 3. Вывести на экран в порядке возрастания три введенных числа
Пример:
Вход: 1, 9, 2
Выход: 1, 2, 9
*/

package main

import "fmt"

func main() {
	var number1, number2, number3, tmp int
	fmt.Println("Введите первое число : ")
	fmt.Scanf("%d\n", &number1)
	fmt.Println("Введите второе число : ")
	fmt.Scanf("%d\n", &number2)
	fmt.Println("Введите третье число : ")
	fmt.Scanf("%d\n", &number3)

	if number1 > number2 {
		tmp = number2
		number2 = number1
		number1 = tmp
	}
	if number1 > number3 {
		tmp = number3
		number3 = number2
		number2 = number1
		number1 = tmp
	} else if number2 > number3 {
		tmp = number3
		number3 = number2
		number2 = tmp
	}
	fmt.Printf("%d, %d, %d\n", number1, number2, number3)

}
