package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	data := generate()
	writeFile(data)
	fmt.Println(findMinMax())
}

func writeFile(data string) {
	file, err := os.Create("file.txt")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(data)

	fmt.Println("Data is in file")
}

func generate() string {
	var result string = ""
	for i := 0; i <= 20; i++ {
		result += strconv.Itoa(rand.Intn(100)) + "\n"
	}
	return result
}

func findMinMax() (int, int) {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	min, max := 100.0, 0.0 // здесь все правильно!
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		num, _ := strconv.ParseFloat(str, 64)
		min = math.Min(float64(min), num)
		max = math.Max(float64(max), num)
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	return int(min), int(max)
}
