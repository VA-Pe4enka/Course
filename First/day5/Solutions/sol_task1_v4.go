package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// var arr []int
	rand.Seed(time.Now().Unix())
	var str string
	var readed []string

	min := math.Inf(1)
	max := math.Inf(-1)

	for i := 0; i < 20; i++ {
		x := rand.Intn(100)
		str += strconv.Itoa(x) + " "
	}

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(str)

	file1, err := os.Open("test.txt")
	scanner := bufio.NewScanner(file1)
	for scanner.Scan() {
		readed = strings.Split(strings.TrimSpace(scanner.Text()), " ")
	}
	err = file1.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	fmt.Println(readed)
	for _, value := range readed {
		x, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(x)
		// fmt.Println(min)
		// fmt.Println(max)
		min = float64(math.Min(min, float64(x)))
		max = float64(math.Max(max, float64(x)))
	}

	fmt.Println(min)
	fmt.Println(max)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
