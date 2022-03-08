package main

import (
        "fmt"
)

func main() {

        var number int

        fmt.Println("enter 2- or 3-digit number:")
        fmt.Scanf("%d", &number)

        if number / 100 < 1 && number > 9{
                result := number / 10 + number % 10
                fmt.Printf("The sum of digits is %d\n", result)
        } else if number/100 < 10 && number/10 > 9 {
                second := (number / 10) % 10
                third := number % 10
                if third == 0 {
                        third = 1
                }
                if second == 0 {
                        second = 1
                }
                result := number / 100 * second * third
                fmt.Printf("The intersection of digits is %d\n", result)
        } else {
                fmt.Printf("Entered unsupported number")
        }
}