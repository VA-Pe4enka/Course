/*
Задача №1
Написать функцию, которая расшифрует строку.
code = "220411112603141304"
Каждые две цифры - это либо буква латинского алфавита в нижнем регистре либо пробел.
Отчет с 00 -> 'a' и до 25 -> 'z', 26 -> ' '(пробел).
Вход: строка из цифр. Выход: Текст.
Проверка работы функции выполняется через вторую строку.

codeToString(code) -> "???????'
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func setAlph() map[byte]rune {
	alph := make(map[byte]rune)
	alph[26] = ' '

	var j byte
	for i := 'a'; i <= 'z'; i++ {
		alph[j] = i
		j++
	}
	return alph
}

func codeToString(code string, alph map[byte]rune) string {
	var resStr string
	for i := 0; i < len(code)-1; i += 2 {
		c, err := strconv.Atoi(code[i : i+2])
		if err != nil {
			fmt.Println("Error")
		}
		resStr += fmt.Sprintf("%c", alph[(byte(c))])
	}
	return resStr
}

func main() {
	//const code = "220411112603141304"
	reader := bufio.NewReader(os.Stdin)
	code, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	code = strings.Trim(code, "\r\n")

	alph := setAlph()

	fmt.Printf("codeToString(%s) -> \"%s\"\n", code, codeToString(code, alph))
}
