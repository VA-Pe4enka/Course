//Домашнее задание. День 3

package main

import (
	"fmt"
	"os"
	"strings"
)

var str string
var attempts = 5

func codeToString(code string) string {

	keys := make(map[string]string)
	keys["00"] = "a"
	keys["01"] = "b"
	keys["02"] = "c"
	keys["03"] = "d"
	keys["04"] = "e"
	keys["05"] = "f"
	keys["06"] = "g"
	keys["07"] = "h"
	keys["08"] = "i"
	keys["09"] = "j"
	keys["10"] = "k"
	keys["11"] = "l"
	keys["12"] = "m"
	keys["13"] = "n"
	keys["14"] = "o"
	keys["15"] = "p"
	keys["16"] = "q"
	keys["17"] = "r"
	keys["18"] = "s"
	keys["19"] = "t"
	keys["20"] = "u"
	keys["21"] = "v"
	keys["22"] = "w"
	keys["23"] = "x"
	keys["24"] = "y"
	keys["25"] = "z"
	keys["26"] = " "

	words := make([]string, len(code)/2)
	j := 0
	for i := 0; i < len(code); i += 2 {
		words[j] = code[i : i+2]
		j++
	}

	for i := 0; i < len(words); i++ {
		for j := range keys {
			if j == words[i] {
				str += keys[j]
			}
		}
	}

	return str
}

func enterPassword() string {
	var password string

	if attempts > 0 {
		fmt.Println("Enter your password. You have ", attempts, " more attempts.")
		fmt.Scanf("%s\n", &password)
		attempts--
	} else {
		fmt.Println("You have no more attempts!!!")
		os.Exit(1)
	}
	return password
}

func checkPassword(password string) {

	const digits = "0123456789"
	const lowercase = "abcdefghiklmnopqrstvxyz"
	const uppercase = "ABCDEFGHIKLMNOPQRSTVXYZ"
	const special = "_!@#$%^&"

	check := 0

	for i := 0; i < attempts; i++ {
		if len(password) < 8 || len(password) > 15 {
			fmt.Println("Invalid password!!! You have ", attempts, " more attempts")
			enterPassword()
		} else {
			if strings.ContainsAny(password, digits) {
				check++
			}

			if strings.ContainsAny(password, lowercase) {
				check++
			}

			if strings.ContainsAny(password, uppercase) {
				check++
			}

			if strings.ContainsAny(password, special) {
				check++
			}

			if check == 4 {
				fmt.Println("That`s a great password!")
				os.Exit(0)
			} else {
				fmt.Println("Password should contain digits, lowercase, uppercase and special symbols! Try again.")
				enterPassword()
			}

		}

	}

}

func main() {

	fmt.Println(codeToString("220411112603141304"))

	enterPassword()
	checkPassword(enterPassword())

}
