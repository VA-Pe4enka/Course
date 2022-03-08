/*
Написать функцию.
Входные аргументы функции: количество бутылок от 0 до 200.
Функция должна вернуть количество и слово бутыл<?> с правильным окончанием.
*/
package main

import (
	"fmt"
	"strconv"
	"log"
)

func main() {
	var amount int

	fmt.Print("Number of bottles (from 0 to 200):\n>>> ")
	_, err := fmt.Scan(&amount)
	if err != nil {
		return
	}

	result, err := ending(amount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func ending(amount int) (string, error) {
	var end, base string
	base = " бутыл"

	if amount < 0 || amount > 200 {
		return "", fmt.Errorf("the number of bottles must be between 0 and 200")
	}

	if amount % 100 < 11 || amount % 100 > 14 {
		switch amount % 10 {
		case 1:
			end = "ка"
		case 2, 3, 4:
			end = "ки"
		default:
			end = "ок"
		}
	} else {
		end = "ок"
	}
	return strconv.Itoa(amount) + base + end, nil
}