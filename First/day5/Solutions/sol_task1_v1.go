/*
С помощью rand.Intn(101) сформировать 20 случайных чисел(от 0 до 100), записать в файл.
Считать эти числа и вывести минимальное и максимальное
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	fileName := "numbers.txt"
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}

	for i := 0; i < 20; i++ {
		n := rand.Intn(100)
		file.WriteString(strconv.Itoa(n) + "\n")
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Write done.")

	file, err = os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var min int = 100
	var max int = 0
	var number int
	for scanner.Scan() {
		number, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if number > max {
			max = number
		}
		if number < min {
			min = number
		}
		fmt.Printf("%v ", number)
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	fmt.Println("Read done.")

}
