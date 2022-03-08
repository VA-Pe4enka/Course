package utils

import (
	"bufio"
	"log"
	"fmt"
	"os"
)

func ReadLine(message string) string {
	var input string
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s> ", message)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(fmt.Sprintf("Произошла ошибка: %s", err))
	}
	return input
}
