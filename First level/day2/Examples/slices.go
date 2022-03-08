package main

import "fmt"

func main() {
	// var notes []string
	// // Создание среза(слайсы)
	// notes = make([]string, 7)
	// notes[4] = "lang Golang"

	// primes := make([]int, 5)
	// primes[2] = 2
	// primes[3] = 89
	// fmt.Println(notes)
	// fmt.Println(primes)
	// letters := []string{"a", "b", "c"}
	// for i := 0; i < len(letters); i++ {
	// 	fmt.Println(letters[i])
	// }
	// for _, letter := range letters {
	// 	fmt.Println(letter)
	// }
	// notes_all := []string{"do", "re", "mi"}
	// primes_all := []int{
	// 	34,
	// 	89,
	// 	12,
	// }
	// fmt.Println(notes_all, primes_all)

	// myArray := [5]string{"a", "b", "c", "d", "e"}

	// //Создание слайса start(0):stop
	// mySlice := myArray[1:3]
	// fmt.Println(mySlice)
	// i, j := 1, 5
	// sliceBig := myArray[i:j]
	// fmt.Println(sliceBig)
	// fmt.Println(myArray[3:])

	// array := [5]string{"do", "re", "mi", "fa", "sol"}
	// sliceOne := array[0:3]
	// sliceTwo := array[2:5]
	// array[2] = "X"
	// fmt.Println(array)
	// fmt.Println(sliceOne, sliceTwo)

	// //Как изменять массив 
	// sliceNew := []string{"a", "b"}
	// fmt.Println(sliceNew, len(sliceNew))

	// // Используем append для изменения слайса
	// sliceNew = append(sliceNew, "c")
	// fmt.Println(sliceNew, len(sliceNew))
	// sliceNew = append(sliceNew, "d", "e")
	// fmt.Println(sliceNew, len(sliceNew))

	s1 := []string{"s1", "s1"}
	s2 := append(s1, "s2", "s2")
	s3 := append(s2, "s3", "s3")
	// нельзя сказать точно, новый slice или базовый
	//s4 := append(s3, "s4", "s4")
	s4 := append(s3, "s4", "s4", "tststhshs")
	s5 := append(s4, "s5", "s5")
	fmt.Println(s1, s2, s3, s4, s5)
	s4[0] = "XYZ"
	fmt.Println(s1, s2, s3, s4, s5)

}