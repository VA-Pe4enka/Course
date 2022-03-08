/* На вход целое число
Если двузначное, на выходе - сумма цифр числа.
Если трехзначное, на выходе - произведение положительных цифр.
Для других чисел вывести, что опереция не может быть выполнена.
*/
package main

import(
	"fmt"
)

func main() {
	var user_num int
	fmt.Println("Enter digit?")
	fmt.Scanf("%d\n", &user_num)

	if user_num < 10 || user_num > 999 {
	    fmt.Println("Операция не может быть выполнена")
	} else if user_num / 100 < 1 {
	    fmt.Println(user_num / 10 + user_num % 10)
	} else if user_num / 100 < 10 {
	    first_digit := user_num / 100
	    second_digit := user_num / 10 % 10
	    third_digit := user_num % 10
	    if second_digit == 0 && third_digit > 0 {
	        fmt.Println(first_digit * third_digit)
	    } else if second_digit > 0 && third_digit == 0 {
	        fmt.Println(first_digit * second_digit)
	    } else {
	        fmt.Println(first_digit * second_digit * third_digit)
	    }
	}
}