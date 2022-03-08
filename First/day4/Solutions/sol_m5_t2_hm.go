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
	"strings"
)

func Cipher(str string) []string {
	cipherMap := map[string]string{
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
	word := make([]string, 0)
	slice := SplitStrings(str, 2)
	for i := 0; i < len(slice); i++ {
		for key, val := range cipherMap {
			if key == slice[i] {
				word = append(word, val)
			}
		}

	}
	return word
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func SplitStrings(s string, chunkLength int) []string {
	length := len(s)

	chunksCount := length / chunkLength
	if length%chunkLength != 0 {
		chunksCount += 1
	}

	chunks := make([]string, chunksCount)

	for i := range chunks {
		from := i * chunkLength
		to := min(length, from+chunkLength)
		chunks[i] = s[from:to]
	}

	return chunks
}

func main() {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	var s string
	fmt.Println(strings.Join(Cipher(line), s))
}
