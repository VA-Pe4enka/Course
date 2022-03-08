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

import "fmt"

func main() {
	var size int
	fmt.Println("Введите размер шахматной доски от 0 до 20:")
	fmt.Scanf("%d\n", &size)

	if size < 0 || size > 20 {
		fmt.Println("Размер задан неверно")
	}

	element := 0
	for i := 0; i < size; i++ {
		var newString string
		for j := 0; j < size; j++ {
			newString = fmt.Sprintf("%s%d ", newString, element)
				element ^= 1
		}
		fmt.Println(newString)
		if size % 2 == 0 {
			element ^= 1
		}
	}
}