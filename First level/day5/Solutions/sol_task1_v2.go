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

func generateSeed(filename string, count int) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Невозможно открыть файл: %s", err)
	}
	defer file.Close()
	for i := 0; i < count; i++ {
		_, err := file.WriteString(fmt.Sprintf("%d\n", rand.Intn(100)))
		if err != nil {
			log.Fatalf("Невозможно записать в файл: %s", err)
		}
	}
}

func readSeed(filename string) []int {
	var (
		result []int
	)
	result = make([]int, 0, 100)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Невозможно открыть файл: %s", err)
	}
	scanner := bufio.NewScanner(file)
	defer file.Close()
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err == nil {
			result = append(result, num)
		}
	}
	return result
}

func getMax(numbers []int) int {
	var max int = numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

func getMin(numbers []int) int {
	var min int = numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	return min
}

func main() {
	rand.Seed(time.Now().Unix())
	var numbers []int
	generateSeed("seed.txt", 20)
	numbers = readSeed("seed.txt")
	fmt.Printf("Min: %d\n", getMin(numbers))
	fmt.Printf("Max: %d\n", getMax(numbers))
}
