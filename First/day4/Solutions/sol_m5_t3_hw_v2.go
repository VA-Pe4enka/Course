/*
Задача №2
Вход:
Пользователь должен ввести правильный пароль, состоящий из:
цифр,
букв латинского алфавита(строчные и прописные) и
специальных символов  special = "_!@#$%^&"

Всего 4 набора различных символов.
В пароле обязательно должен быть хотя бы один символ из каждого набора.
Длина пароля от 8(мин) до 15(макс) символов.
Максимальное количество попыток ввода неправильного пароля - 5.
Каждый раз выводим номер попытки.
*Желательно выводить пояснение, почему пароль не принят и что нужно исправить.

digits = "0123456789"
lowercase = "abcdefghijklmnopqrstuvwxyz"
uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
special = "_!@#$%^&"

Выход:
Написать, что ввели правильный пароль.

Пример:
хороший пароль -> o58anuahaunH!
хороший пароль -> aaaAAA111!!!
плохой пароль -> saucacAusacu8
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	var password string

	i := 0
	for i < 5 {
		i++
		password = (readLine(fmt.Sprintf("Попытка %d. Введите пароль", i)))
		if checkPassword(password) {
			fmt.Println("Хороший пароль!")
			break
		}

	}
}

func checkPassword(password string) bool {

	if len(password) < 8 || len(password) > 15 {
		fmt.Println("Длина пароля должна быть от 8(мин) до 15(макс) символов.")
		return false
	}

	digits := "0123456789"
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special := "_!@#$%^&"

	digitsCount := 0
	lowercaseCount := 0
	uppercaseCount := 0
	specialCount := 0

	for _, chr := range password {
		if strings.Contains(digits, string(chr)) {
			digitsCount++
		} else if strings.Contains(lowercase, string(chr)) {
			lowercaseCount++
		} else if strings.Contains(uppercase, string(chr)) {
			uppercaseCount++
		} else if strings.Contains(special, string(chr)) {
			specialCount++
		} else {
			fmt.Printf("Недопустимый символ '%c'\n", chr)
			return false
		}
	}

	if digitsCount == 0 {
		fmt.Println("Пароль должен содержать цифры")
		return false
	}
	if lowercaseCount == 0 {
		fmt.Println("Пароль должен содержать маленькие буквы")
		return false
	}
	if uppercaseCount == 0 {
		fmt.Println("Пароль должен содержать большие буквы")
		return false
	}
	if specialCount == 0 {
		fmt.Println("Пароль должен содержать спецсимволы _!@#$%^&")
		return false
	}

	return true
}
