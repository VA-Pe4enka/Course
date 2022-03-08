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
	var inCode string
	var decodedString string

	fmt.Println("Штирлиц, Я жду шифр")

	decoder := createEngABC()
	inCode = ask()

	if len(inCode)%2 == 0 {
		for i := 0; i < len(inCode); i += 2 {
			if idx, err := strconv.Atoi(inCode[i : i+2]); err == nil {
				s := decoder[idx]
				decodedString += s
			} else if err != nil {
				log.Fatal(err)
			}

		}
		fmt.Println(decodedString)
	} else {
		fmt.Println("Шифр не верный! Вы предатель!")
	}

}

func ask() string {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	data := strings.TrimSpace(line)
	return data
}

func createEngABC() map[int]string {
	abc := "abcdefghijklmnopqrstuvwxyz "
	m := make(map[int]string, len(abc))
	for i := 0; i < len(abc); i++ {
		m[i] = abc[i : i+1]
	}
	return m
}
