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
)

const (
	digits = "0123456789"
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special = "_!@#$%^&"
)

func setAlph (str string) map[string]int {
	mAlph := make(map[string]int, len(str))
	for _, char := range str {
		mAlph[string(char)] = 0
	}
	return mAlph
}

func cross(pswd string, mChars map[string]int) bool {
	for i := 0; i < len(pswd); i++ {
		_, ok := mChars[pswd[i:i+1]]
		if ok {
			return true
		}
	}
	return false
}

func checkPswd(pswd string) bool {
	mapDigits := setAlph(digits)
	mapLowercase := setAlph(lowercase)
	mapUppercase := setAlph(uppercase)
	mapSpecial := setAlph(special)

	var flag uint8
	if !cross(pswd, mapDigits) {
		fmt.Printf("В пароле не хватает %s\n", digits)
	} else {
		flag++
	}

	if !cross(pswd, mapLowercase) {
		fmt.Printf("В пароле не хватает %s\n", lowercase)
	} else {
		flag++
	}

	if !cross(pswd, mapUppercase) {
		fmt.Printf("В пароле не хватает %s\n", uppercase)
	} else {
		flag++
	}

	if !cross(pswd, mapSpecial) {
		fmt.Printf("В пароле не хватает %s\n", special)
	} else {
		flag++
	}

	if flag == 4 {
		fmt.Println("Прекрасный пароль!")
		return true
	} else {
		return false
	}
}

func pswdEnter() string {
	fmt.Println("Введите пароль: ")
	reader := bufio.NewReader(os.Stdin)
	pswd, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return pswd
}

func checkPswdLen(pswdLen int) bool {
	if  pswdLen >= 8 && pswdLen <= 15 {
		return false
	}	
	fmt.Println("Длина пароля должна быть от 8 до 15 символов.")
	return true
}

func main() {
	var pswd string
	for i := 5; i > 0; i-- {
		flagLen := true
		for flagLen {
			pswd = pswdEnter()
			flagLen = checkPswdLen(len(pswd))
		}
		if !checkPswd(pswd) {
			fmt.Printf("Попробуйте ввести пароль еще раз (: \n")
			fmt.Printf("Осталось %d попыток. \n", i-1)
		} else {
			fmt.Println("Пароль принят!")
			break
		}
	}
}