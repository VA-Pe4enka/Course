package main

import (
	"fmt"
	"time"
)

func main(){
	var notes [7]string
	var primes [4]int
	var dates [3]time.Time
	var numbers [5]int = [5]int{8, 12, 7, 8, 12}
	strs := [7]string{"do", "re", "mi", "fa", "so", "la", "si"}
	text := [3]string{
		"this is first line",
		"it's second line",
		"third line",
	}
	notes[0] = "foo"
	notes[1] = "bar"
	notes[2] = "baz"
	primes[3] = 56
	primes[0] = 1
	primes[1]++
	primes[2] += 2

	dates[0] = time.Unix(1257894000, 0)
	dates[1] = time.Unix(1447920000, 0)

	fmt.Println(notes[0], notes[1])
	fmt.Println(primes)
	fmt.Println(primes[0], primes[1])
	fmt.Println(dates[1], dates[0])
	fmt.Println("numbers[0] = ", numbers[0])
	fmt.Println(strs)
	fmt.Println("hello, world")
	fmt.Println(text)
	fmt.Printf("%#v\n", notes)
	fmt.Printf("%#v\n", text)
	fmt.Printf("%#v\n", numbers)

	// for i := 0; i < len(strs); i++ {
	// 	fmt.Println(i, strs[i])
	// }
	// for index, value := range text {
	// 	fmt.Println(index, value)
	// }
}