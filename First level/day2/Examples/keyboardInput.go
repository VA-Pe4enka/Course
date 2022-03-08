package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
	"log"
)

//GetFloat считывает значение со стандартного ввода
//Возвращает само число и ошибку.
func GetFloat() (float64, error) {
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
	return number, nil
}

func main() {
 	result, err := GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}