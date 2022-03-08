/*
Написать функцию, которая расшифрует строку.
Каждые две цифры - это либо буква латинского алфавита в нижнем регистре либо пробел.
Отчет с 00 -> 'a' и до 25 -> 'z', 26 -> ' '(пробел).
Вход: строка из цифр. Выход: Текст.
Проверка работы функции выполняется через вторую строку.
*/
package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	var code string
	fmt.Print("Cipher for decoding:\n>>> ")
	_, err := fmt.Scan(&code) // example: 220411112603141304
	check(err)
	fmt.Println(reverseString(codeToString(code)))
}

func codeToString(code string) []string {
	list := make([]string, len(code)/2)
	num, err := strconv.Atoi(code)
	check(err)
	for i := 0; i < len(code)/2; i++ {
		a := num % 100
		if a > 26 {
			log.Fatalf("wrong data, %d is an invalid value\n", a)
		} else if a == 26 {
			list[i] = " "
		} else {
			list[i] = string(rune(97 + a))
		}
		num /= 100
	}
	return list
}

func reverseString(list []string) string {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
	return strings.Join(list, "")
}

func check(err error) {
	if err != nil {
		log.Fatalln("wrong data")
	}
}
