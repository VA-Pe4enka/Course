/*
Написать функцию.
Входные аргументы функции: количество бутылок от 0 до 200.
Функция должна вернуть количество и слово бутыл<?> с правильным окончанием.
Пример :
In: 5
Out: 5 бутылок

In: 1
Out: 1 бутылка

In: 22
Out: 22 бутылки
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

func main() {
	fmt.Println("Введите число 0-200")
	result, err := GetInt()
	if err != nil {
		log.Fatal(err)
	}

	if result < 0 || result > 200 {
		fmt.Println("Число не находится между 0 и 200")
	} else {
		fmt.Println(WordForm(result))
	}
}

func GetInt() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return int(number), nil
}
func WordForm(number int) string {

	var forms = [3]string{"ка", "ки", "ок"}
	var array = [6]int{2, 0, 1, 1, 1, 2}
	var hund = number % 100
	var dec = number % 10
	// 5-19 true, остальные false
	inHund := hund > 4 && hund < 20
	inDec := dec < 5

	form := array[5] // form = 2
	if !inHund && inDec {
		form = array[dec]
	}
	return "бутыл" + forms[form]
}
