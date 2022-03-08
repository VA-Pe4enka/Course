/* Вход: строка
   Выход: частота для каждого символа
   fer example:
   hello go-> h: 1, e: 1, l:2, o:2, g: 1, " ":1
   strings.Split, strings.Fields
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func countChars(message string) map[string]int {
	count := make(map[string]int)

	for _, char := range message {
		count[string(char)]++
	}
	return count
}

func readLine(message string) string {
	var input string
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s> ", message)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(fmt.Sprintf("Error: %s", err))
	}
	return input
}

func main() {
	var (
		message string
		counter map[string]int
	)

	message = strings.Trim(strings.ToLower(readLine("Введите текст")), "\n")

	counter = countChars(message)

	for key, value := range counter {
		fmt.Printf("%s: %d\n", key, value)
	}
}
