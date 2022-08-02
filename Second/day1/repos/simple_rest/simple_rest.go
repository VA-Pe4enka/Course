package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Operation struct {
	Type   string `json:"Type"`
	Result int    `json:"Result"`
}

type Message struct {
	Message string `json:"Message"`
	Payload []int  `json:"Payload,omitempty"`
}

const randomCount = 100
const randomsFile = "/tmp/randoms.txt"

func seedExists() bool {
	_, err := os.Stat(randomsFile)
	return err == nil
}

func createRandoms() {
	if seedExists() {
		return
	}
	rand.Seed(time.Now().Unix())
	file, err := os.Create(randomsFile)
	if err != nil {
		log.Fatalf("Unable to create randoms file: %s", err)
	}
	defer file.Close()
	for i := 0; i < randomCount; i++ {
		_, err := file.WriteString(fmt.Sprintf("%d\n", rand.Intn(100)))
		if err != nil {
			log.Fatalf("Unable to write to randoms file: %s", err)
		}
	}
}

func readRandoms() []int {
	var result []int
	result = make([]int, 0, randomCount)
	file, err := os.Open(randomsFile)
	if err != nil {
		log.Fatalf("Unable to read randoms file: %s", err)
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

func sum(iterable []int) int {
	var sum = 0
	for _, value := range iterable {
		sum += value
	}
	return sum
}

func FirstHandler(w http.ResponseWriter, r *http.Request) {
	createRandoms()
	randoms := readRandoms()

	operation := Operation{
		Type:   "first",
		Result: randoms[0],
	}

	json.NewEncoder(w).Encode(operation)
}

func SecondHandler(w http.ResponseWriter, r *http.Request) {
	createRandoms()
	randoms := readRandoms()

	operation := Operation{
		Type:   "second",
		Result: randoms[1],
	}

	json.NewEncoder(w).Encode(operation)
}

func SumHandler(w http.ResponseWriter, r *http.Request) {
	createRandoms()
	randoms := readRandoms()

	operation := Operation{
		Type:   "sum",
		Result: sum(randoms),
	}

	json.NewEncoder(w).Encode(operation)
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	createRandoms()
	randoms := readRandoms()

	json.NewEncoder(w).Encode(Message{
		Message: "Current random set is in Payload",
		Payload: randoms,
	})
}

func main() {
	http.HandleFunc("/", MainHandler)
	http.HandleFunc("/first", FirstHandler)
	http.HandleFunc("/second", SecondHandler)
	http.HandleFunc("/sum", SumHandler)
	http.ListenAndServe(":1234", nil)
}
