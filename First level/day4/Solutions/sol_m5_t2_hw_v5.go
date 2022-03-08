package main

import (
	"fmt"
	"log"
	"strings"
	"example/utils"
)

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

func invertMap(sourcemap map[string]string) map[string]string {
	result := make(map[string]string)

	for key, value := range sourcemap {
		result[value] = key
}

	return result
}

func generateVocabulary(start int) map[string]string {
	result := make(map[string]string)

	for i := start; i < 26; i++ {
		result[fmt.Sprintf("%c", 'a'+i)] = fmt.Sprintf("%02d", start+i)
	}

	result[" "] = "26"

	return result
}

func encode(message string) string {
	var (
		vocabulary map[string]string = generateVocabulary(0)
		result     string            = ""
	)

	for _, char := range strings.Trim(strings.ToLower(message), " \n") {
		value, ok := vocabulary[string(char)]
		if !ok {
			continue
		} // Пропускаем символ, если его нет в словаре
		//if !ok { log.Fatalf("Неопознанный символ: %c", char)	} // Прерываем выполнение, если символа нет в словаре
		result += value

	}

	return result
}

func decode(message string) string {
	var (
		vocabulary map[string]string = invertMap(generateVocabulary(0))
		result     string            = ""
	)

	for i := 0; i < len(message)-1; i += 2 {
		char := message[i : i+2]
		value, ok := vocabulary[char]
		if !ok {
			log.Fatalf("Неопознанная последовательность: %s", char)
		}
		result += value
	}

	return result
}

func main() {
	var choice string

	choice = strings.ToLower(strings.Trim(utils.ReadLine("[E] Кодировать, [D] Декодировать"), " \n"))

	switch choice {
	case "e":
		fmt.Println(encode(utils.ReadLine("Введите сообщение")))
	case "d":
		fmt.Println(decode(utils.ReadLine("Введите шифр")))
	default:
		log.Fatalf("Неправильный выбор: %s", choice)
	}
}