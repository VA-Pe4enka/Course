package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func saveToJSON(filename *os.File, key interface{}) {
	encodeJSON := json.NewEncoder(filename)
	err := encodeJSON.Encode(key)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	myRecord := Record{
		Name:    "Михаил",
		Surname: "Иванов",
		Tel: []Telephone{
			Telephone{Mobile: true, Number: "1234-567"},
			Telephone{Mobile: true, Number: "1234-abcd"},
			Telephone{Mobile: false, Number: "abcc-567"},
		},
	}

	f, err := os.Create("sample.json")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	saveToJSON(os.Stdout, myRecord)
	saveToJSON(f, myRecord)

}
