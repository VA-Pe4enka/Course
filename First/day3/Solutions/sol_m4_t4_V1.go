/*
Задача №4. Шахматная доска
Вход: размер шахматной доски, от 0 до 20
Выход: вывести на экран эту доску, заполняя поля нулями и единицами

Пример:
Вход: 5
Выход:
    0 1 0 1 0
    1 0 1 0 1
    0 1 0 1 0
    1 0 1 0 1
    0 1 0 1 0
*/
package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Println("Введите размер шахматной доски, от 0 до 20:")
	fmt.Scanf("%d\n", &number)
	if number < 0 || number > 20 {
		fmt.Println("Неверно задано значение")
		return

	}
	var firstSing bool = false
	var sing bool
	var temp int = 1
	for i := 0; i < number; i++ {
		sing = firstSing

		for j := 0; j < number; j++ {
			fmt.Printf("%d ", boolToInt(sing))
			sing = !sing
		}

		if firstSing == sing {
			firstSing = !sing
		} else {
			firstSing = sing
		}

		fmt.Println()
	}
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
