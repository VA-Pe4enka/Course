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
	"strconv"
)

func main() {
	var number int
	var line string
	fmt.Println("Введите размер шахматной доски, от 0 до 20 : ")
	fmt.Scanf("%d\n", &number)
	if 0 <= number && number <= 20 {
		for i := 0; i < number; i++ {
			line = ""
			for j := 0; j < number; j++ {
				line += strconv.Itoa((i + j) % 2)
				line += " "
			}
			fmt.Println(line)
		}

	} else {
		fmt.Println("Введены неверные данные.")
	}
}
