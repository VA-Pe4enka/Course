package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

//Guestbook - структура, используемая при отображении view.html
type Guestbook struct {
	SignatureCount int
	Signatures     []string
}

//check вызывает log.Fatal для любых ошибок, отличных от nil
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//getStrings возращает срез строк, прочитанный из fileName,
//по одной строке на каждую строку файла.
func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines
}

//viewHandler читает записи гостевой книги и
//выводит их вместе со счетчиком
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("signatures.txt")
	html, err := template.ParseFiles("view.html")
	check(err)
	guestbook := Guestbook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
	err = html.Execute(writer, guestbook)
	check(err)
}

//newHander отображает форму для ввода записи
func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("new.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

//createHandler отображает форму для ввода записи
func createHandler(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature")
	//_, err := writer.Write([]byte(signature)) // only for testing
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	check(err)
	_, err = fmt.Fprintln(file, signature)
	check(err)
	err = file.Close()
	check(err)
}

func main() {

	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
