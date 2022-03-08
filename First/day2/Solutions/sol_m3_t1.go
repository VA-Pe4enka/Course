/*
Вход: слово(без пробела)
Выход: количество буквы 'a'(латинская) в слове. Если нет, то написать об этом.
*/
package main

import "fmt"

func main() {
	var str string
	const character = 'a'
	var count int

	fmt.Println("Введите слово : ")
	fmt.Scanf("%s\n", &str)

	for _, value := range str {
		if value == character {
			count++
		}
	}
	if count != 0 {
		fmt.Printf("Буква 'a' встречается в слове '%s' %d раз(-а)\n", str, count)
	} else {
		fmt.Printf("Буква 'a' не встречается в слове '%s'\n", str)
	}
}
