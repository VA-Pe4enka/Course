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

func readLine(message string) string {
	var input string
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s> ", message)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(fmt.Sprintf("Error: %s", err))
	}
	input = strings.Trim(strings.Trim(input, "\n"), "\r")

	return input
}

func main() {

	code := "220411112603141304"

	fmt.Print("220411112603141304 -> ")
	fmt.Println(codeToString(code))

	str := readLine("Введите код для проверки")
	if checkString(str) {
		fmt.Print(str + " -> ")
		fmt.Println(codeToString(str))
	} else {
		fmt.Println("Код должен состоять из четного числа цифр")
	}

}

func codeToString(code string) string {
	var abcMap map[string]rune = make(map[string]rune)
	var a int = int('a')
	for i := 0; i < 26; i++ {
		abcMap[fmt.Sprintf("%02d", i)] = rune(a + i)
	}
	abcMap["26"] = ' '

	//fmt.Printf("%v", abcMap)

	result := ""
	for j := 0; j < len(code); j += 2 {
		result += string(abcMap[string(code[j:j+2])])
	}

	return result
}

func checkString(code string) (flag bool) {
	flag = false
	if len(code)%2 != 0 {
		return
	}
	for j := 0; j < len(code); j++ {
		_, err := strconv.Atoi(string(code[j]))
		if err != nil {
			return
		}
	}
	flag = true
	return
}
