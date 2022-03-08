package main
 
import (
    "fmt"
)
 
// Задача 2
 
 
func main() {
    var integer int
 
    fmt.Println("Введите трехзначное число")
    fmt.Printf("> ")
    fmt.Scanf("%d", &integer)
	
	// && - Логическое "И", || - Логическое "ИЛИ"
    if 99 < integer && integer < 1000 {
        var (
            first int = integer / 100
            last int = integer % 10
        )
        fmt.Printf("Первое число: %d, последнее: %d\n", first, last)
    } else {
        fmt.Println("Число не является трехзначным")
    }
	//true + false -> true
	if integer / 100 > 7 || integer % 10 > 3 {
		fmt.Println("Проверка выполнена")
	}
	flag := false
	if integer > 0 {
		flag = true
		fmt.Println("Number is positive")
	} else if integer < 0 {
		fmt.Println("Number is negative")
	} else {
		fmt.Println("Number is equal zero")
	}
	fmt.Println(flag)
}
