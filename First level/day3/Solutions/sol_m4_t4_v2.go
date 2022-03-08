package main

import (
	"fmt"
)

func main() {
	var chessBoardSize int

	fmt.Println("введите размер шахматной доски, от 0 до 20")
	fmt.Scan(&chessBoardSize)

	if chessBoardSize >= 0 && chessBoardSize <= 20 {

		for i := 0; i < chessBoardSize; i++ {
			if i%2 == 0 {
				for i := 0; i < chessBoardSize; i++ {
					if i%2 == 0 {
						fmt.Print("0 ")
					} else {
						fmt.Print("1 ")
					}
				}
			} else {
				for i := 0; i < chessBoardSize; i++ {
					if i%2 != 0 {
						fmt.Print("0 ")
					} else {
						fmt.Print("1 ")
					}
				}
			}
			fmt.Println()
		}
	} else {
		fmt.Println("сказал же, от 0 до 20")
	}
}