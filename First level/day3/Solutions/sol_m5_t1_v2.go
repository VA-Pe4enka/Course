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

func main() {
	fmt.Println("введите фразу (латиница)")
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	line = strings.Trim(line, "\n")

	elements := make(map[rune]int)
	for _, value := range line {
		elements[value]++
	}

	for k, v := range elements {
		fmt.Printf("Key: %v, value: %v\n", string(k), v)
	}
}