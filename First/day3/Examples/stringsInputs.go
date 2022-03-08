package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	data := strings.Fields(line)
	fmt.Println(data)
}
