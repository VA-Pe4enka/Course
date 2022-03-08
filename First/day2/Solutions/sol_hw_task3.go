/*
Задача № 3. Вывести на экран в порядке возрастания три введенных числа
Пример:
Вход: 1, 9, 2
Выход: 1, 2, 9
*/

package main

import "fmt"

func main() {
	var num1, num2, num3 int

	fmt.Println("Введите первое число:")
	fmt.Scanf("%d\n", &num1)

	fmt.Println("Введите второе число:")
	fmt.Scanf("%d\n", &num2)

	fmt.Println("Введите третье число:")
	fmt.Scanf("%d\n", &num3)

	if num1 < num2 && num1 < num3 {
		fmt.Println(num1)
		if num2 < num3 {
			fmt.Println(num2)
			fmt.Println(num3)
		} else {
			fmt.Println(num3)
			fmt.Println(num2)
		}
	} else if num2 < num3 {
		fmt.Println(num2)
		if num1 < num3 {
			fmt.Println(num1)
			fmt.Println(num3)
		} else {
			fmt.Println(num3)
			fmt.Println(num1)
		}
	} else {
		fmt.Println(num3)
		if num2 < num1 {
			fmt.Println(num2)
			fmt.Println(num1)
		} else {
			fmt.Println(num1)
			fmt.Println(num2)
		}
	}
}
