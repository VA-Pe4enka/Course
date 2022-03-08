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

import "fmt"

func main() {
	alphabet := map[string]string{
		"00": "a",
		"01": "b",
		"02": "c",
		"03": "d",
		"04": "e",
		"05": "f",
		"06": "g",
		"07": "h",
		"08": "i",
		"09": "j",
		"10": "k",
		"11": "l",
		"12": "m",
		"13": "n",
		"14": "o",
		"15": "p",
		"16": "q",
		"17": "r",
		"18": "s",
		"19": "t",
		"20": "u",
		"21": "v",
		"22": "w",
		"23": "x",
		"24": "y",
		"25": "z",
		"26": " ",
	}

	var codeString string
	fmt.Println("Введите закодированную строку")
	fmt.Scan(&codeString)
	//codeString = "220411112603141304"
	codeSlice := []string{}

	for i := 0; i < len(codeString)/2; i++ {
		codeSlice = append(codeSlice, codeString[2*i:2*i+2])
	}

	var resultString string

	for _, code := range codeSlice {
		for key, value := range alphabet {

			if code == key {
				resultString = resultString + value
			}
		}
	}
	fmt.Println(resultString)
}
