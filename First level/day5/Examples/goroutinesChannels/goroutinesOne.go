package main

import (
	"fmt"
	"time"
)

func a() {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
}
func b() {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
}

func main() {
	// Запуск горутин
	go a()
	go b()
	//Пример анонимной функции-горутины
	go func(x int) {
		for i := 0; i < x; i++ {
			fmt.Print("c")
		}
	}(42)
	time.Sleep(time.Second)
	fmt.Println(" end main()")
}
