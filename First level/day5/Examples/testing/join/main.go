package main

import (
	"fmt"
	"prose"
)

func main() {
	phrases := []string{"друзья", "недруги"}
	fmt.Println("Вместе:", prose.JoinWithCommas(phrases))
	phrases = []string{"первый", "второй", "третий"}
	fmt.Println("Вместе:", prose.JoinWithCommas(phrases))
}
