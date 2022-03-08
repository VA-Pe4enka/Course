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
)

func main() {
	var howManyBottles int

	fmt.Println("введите сколько бутылок (от 0 до 200)")
	fmt.Scan(&howManyBottles)

	if 0 <= howManyBottles && howManyBottles <= 200 {

		t := howManyBottles % 100

		if 0 <= t && t <= 20 {
			switch t {
			case 2, 3, 4:
				fmt.Printf("%v бутылки\n", howManyBottles)
			case 1:
				fmt.Printf("%v бутылка\n", howManyBottles)
			default:
				fmt.Printf("%v бутылок\n", howManyBottles)
			}
		} else {
			switch t % 10 {
			case 2, 3, 4:
				fmt.Printf("%v бутылки\n", howManyBottles)
			case 1:
				fmt.Printf("%v бутылка\n", howManyBottles)
			default:
				fmt.Printf("%v бутылок\n", howManyBottles)
			}
		}

	} else {
		fmt.Printf("число %v не входит в диапазон запрошенных бутылок\n", howManyBottles)
	}

}
