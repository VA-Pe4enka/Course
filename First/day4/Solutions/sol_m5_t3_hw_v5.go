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
lowercase = "abcdefghiklmnopqrstvxyz"
uppercase = "ABCDEFGHIKLMNOPQRSTVXYZ"
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

func getString() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return input, err
	}

	input = strings.TrimSpace(input)
	return input, nil
}

func checkPassword(password string) (bool, string) {
	err := "Введен неправильный пароль:"
	resLen, errLen := checkLength(password)
	resSpec := checkUseRequiredCharacters(password, "_!@#$%^&")
	resDigit := checkUseRequiredCharacters(password, "0123456789")
	resUpper := checkUseRequiredCharacters(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	resLower := checkUseRequiredCharacters(password, "abcdefghijklmnopqrstuvwxyz")
	if resLen && resSpec && resDigit && resUpper && resLower {
		return true, ""
	}
	if !resLen {
		err += "\n\t"
		err += errLen
	}
	if !resSpec {
		err += "\n\tотсутствуют спецсимволы"
	}
	if !resDigit {
		err += "\n\tотсутствуют цифры"
	}
	if !resUpper {
		err += "\n\tотсутствуют буквы в верхнем регистре"
	}
	if !resLower {
		err += "\n\tотсутствуют буквы в нижнем регистре"
	}
	return false, err

}

func checkLength(password string) (bool, string) {
	if len(password) < 8 {
		return false, "длина недостаточна (менее 8 символов)"
	}

	if len(password) > 15 {
		return false, "длина избыточна (более 15 символов)"
	}
	return true, ""
}

func checkUseRequiredCharacters(password, symbols string) bool {
	for _, charP := range password {
		for _, charS := range symbols {
			if charP == charS {
				return true
			}
		}
	}
	return false
}

func main() {
	flCorrectPassword := false
	for i := 1; i < 6 && !flCorrectPassword; i++ {
		fmt.Printf("Введите пароль (попытка %d): ", i)

		result, err := getString()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
		res, err_msg := checkPassword(result)
		if !res {
			fmt.Println(err_msg)
		} else {
			flCorrectPassword = true
		}
	}
	if flCorrectPassword {
		fmt.Println("Введен хороший пароль")
	} else {
		fmt.Println("Превышено число допустимых попыток")
	}
}
