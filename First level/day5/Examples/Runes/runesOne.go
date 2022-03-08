package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	asciiString := "ABCDE"
	utf8String := "АБСДЕ"

	//len выводит количеств байт, а не количество символов
	fmt.Println(len(asciiString))
	fmt.Println(len(utf8String))
	
	//RuneCountInString выводит количество символов
	fmt.Println(utf8.RuneCountInString(asciiString))
	fmt.Println(utf8.RuneCountInString(utf8String))
}