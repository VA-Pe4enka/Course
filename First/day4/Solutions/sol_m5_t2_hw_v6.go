/*
Задача №1
Написать функцию, которая расшифрует строку.
code = "220411112603141304"
Каждые две цифры - это либо буква латинского алфавита в нижнем регистре либо пробел.
Отчет с 00 -> 'a' и до 25 -> 'z', 26 -> ' '(пробел).
Вход: строка из цифр. Выход: Текст.
Проверка работы функции выполняется через вторую строку.

codeToString(code) -> "???????'
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
	fmt.Println("Введите строку : ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return input, err
	}

	input = strings.TrimSpace(input)
	return input, nil
}

func codeToString(code string) string {
	ranks := make(map[string]string)
	ranks["00"] = "a"
	ranks["01"] = "b"
	ranks["02"] = "c"
	ranks["03"] = "d"
	ranks["04"] = "e"
	ranks["05"] = "f"
	ranks["06"] = "g"
	ranks["07"] = "h"
	ranks["08"] = "i"
	ranks["09"] = "j"
	ranks["10"] = "k"
	ranks["11"] = "l"
	ranks["12"] = "m"
	ranks["13"] = "n"
	ranks["14"] = "o"
	ranks["15"] = "p"
	ranks["16"] = "q"
	ranks["17"] = "r"
	ranks["18"] = "s"
	ranks["19"] = "t"
	ranks["20"] = "u"
	ranks["21"] = "v"
	ranks["22"] = "w"
	ranks["23"] = "x"
	ranks["24"] = "y"
	ranks["25"] = "z"
	ranks["26"] = " "
	var outstr string
	for i := 0; i < len(code); i += 2 {
		outstr += ranks[code[i:i+2]] //0:2, 2:4, 4:6
	}
	return outstr
}

func main() {
	result, err := getString()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(codeToString(result))
}
