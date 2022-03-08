/*
Задача №4. Шахматная доска
Вход: размер шахматной доски, от 0 до 20
Выход: вывести на экран эту доску, заполняя поля нулями и единицами
*/
package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	var boardSize int

	fmt.Print("Chess board dimension (from 0 to 20):\n>>> ")
	_, err := fmt.Scan(&boardSize)
	if err != nil {
		log.Fatal("wrong data")
	}
	if boardSize < 0 || boardSize > 20 {
		log.Fatal("Chess board dimension must be between 0 and 20")
	}

	board := make([][]string, boardSize)

	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if i%2 != j%2 {
				board[i] = append(board[i], "▓▓▓")
			} else {
				board[i] = append(board[i], "░░░")
			}
		}
	}

	for _, subslice := range board {
		fmt.Println(strings.Join(subslice, ""))
	}
}
