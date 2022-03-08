/*
Написать функцию.
Входные аргументы функции: количество бутылок от 0 до 200.
Функция должна вернуть количество и слово бутыл<?> с правильным окончанием.
Пример :
In: 5
Out: 5 бутылок

In: 1
Out: 1 бутылка

In: 22
Out: 22 бутылки
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int
	fmt.Println("Введите количество бутылок от 0 до 200:")
	fmt.Scanf("%d\n", &number)

	fmt.Println(bottels(number))
}

func bottels(number int) string {

	str := " бутыл"

	if number < 0 || number > 200 {
		return "Введено неверное число"
	}
	number %= 100
	if number == 0 || number >= 5 && number <= 20 {
		str += "ок"
	} else if number%10 == 1 {
		str += "ка"
	} else if number%10 >= 2 && number%10 <= 4 {
		str += "ки"
	}

	return strconv.Itoa(number) + str

}
