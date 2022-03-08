package utils

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func Ask(message string, format string, ptr ...interface{}) {
	fmt.Println(message)
	fmt.Printf("> ")
	_, err := fmt.Scanf(format, ptr...)
	if err != nil {
		log.Fatal("Неверный формат")
	}
}

func ReverseInt(integer int) int {
	var result int = 0

	if 999 < integer && integer < 10000 {
		result += (integer % 10) * 1000
		result += ((integer / 10) % 10) * 100
		result += ((integer / 100) % 10) * 10
		result += integer / 1000
	} else if 99 < integer && integer < 1000 {
		result += (integer % 10) * 100
		result += ((integer / 10) % 10) * 10
		result += integer / 100
	} else {
		log.Fatal("Число вне диапазона")
	}

	return result
}

func Pluralize(count int, variants []string) string {
	pluralCases := [6]int{2, 0, 1, 1, 1, 2}

	var variant int

	if count%100 > 4 && count%100 < 20 {
		variant = 2
	} else {
		variant = pluralCases[int(math.Min(float64(count%10), 5))]
	}
	return fmt.Sprintf("%d %s", count, variants[variant])
}

func ReadLine(message string) string {
	var input string
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s> ", message)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(fmt.Sprintf("Произошла ошибка: %s", err))
	}
	return strings.Trim(input, "\n\r")
}
