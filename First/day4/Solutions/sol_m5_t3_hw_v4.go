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

import "fmt"

func main() {
	const tryCounts int = 5
	const success int = 5

	var tryPassword string

	for i := 1; i <= tryCounts; i++ {
		fmt.Printf("Введите пароль. Попытка %d из 5\n", i)
		fmt.Scan(&tryPassword)

		checkNumber := 0

		lenght := checkLenght(tryPassword)
		if !lenght {
			fmt.Println("Некорректная длина. Длина пароля должна быть от 8 до 15 символов")
		} else {
			checkNumber++
		}

		didgits := checkDidgits(tryPassword)
		if !didgits {
			fmt.Println("Пароль должен содержать хотя бы одну цифру")
		} else {
			checkNumber++
		}

		lowercase := checkLowercase(tryPassword)
		if !lowercase {
			fmt.Println("Пароль должен содержать хотя бы одну букву в нижнем регистре")
		} else {
			checkNumber++
		}

		uppercase := checkUppercase(tryPassword)
		if !uppercase {
			fmt.Println("Пароль должен содержать хотя бы одну букву в верхнем регистре")
		} else {
			checkNumber++
		}

		specialSymbols := checkSpecialSymbols(tryPassword)
		if !specialSymbols {
			fmt.Println("Пароль должен содержать хотя бы один спецсимвол (_!@#$%^&)")
		} else {
			checkNumber++
		}

		if checkNumber == success {
			fmt.Println("Пароль принят")
			break
		}
	}
}

func checkLenght(tryPassword string) bool {
	if len(tryPassword) > 7 && len(tryPassword) < 16 {
		return true
	} else {
		return false
	}
}

func checkDidgits(tryPassword string) bool {
	digits := "0123456789"

	for _, pass := range tryPassword {
		for _, didgit := range digits {
			if pass == didgit {
				return true
			}
		}
	}
	return false
}

func checkLowercase(tryPassword string) bool {
	lowercase := "abcdefghijklmnopqrstuvwxyz"

	for _, pass := range tryPassword {
		for _, letter := range lowercase {
			if pass == letter {
				return true
			}
		}
	}
	return false
}

func checkUppercase(tryPassword string) bool {
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for _, pass := range tryPassword {
		for _, letter := range uppercase {
			if pass == letter {
				return true
			}
		}
	}
	return false
}

func checkSpecialSymbols(tryPassword string) bool {
	special := "_!@#$%^&"

	for _, pass := range tryPassword {
		for _, symbol := range special {
			if pass == symbol {
				return true
			}
		}
	}
	return false
}
