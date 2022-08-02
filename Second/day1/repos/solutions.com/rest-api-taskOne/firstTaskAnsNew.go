package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type number struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
}

type numbers struct {
	First  number
	Second number
	Summa  number
}

func (ctx *numbers) FirstHandler(w http.ResponseWriter, r *http.Request) {
	ctx.First = number{"First", ctx.First.Number}
	json.NewEncoder(w).Encode(ctx.First)
}

func (ctx *numbers) SecondHandler(w http.ResponseWriter, r *http.Request) {
	ctx.Second = number{"Second", ctx.Second.Number}
	json.NewEncoder(w).Encode(ctx.Second)
}

func (ctx *numbers) SummaHandler(w http.ResponseWriter, r *http.Request) {
	ctx.Summa = number{"summa", ctx.Summa.Number}
	json.NewEncoder(w).Encode(ctx.Summa)
}

func initCtx() *numbers {
	var ctx numbers
	ctx.First.Number = rand.Intn(100)
	ctx.Second.Number = rand.Intn(100)
	ctx.Summa.Number = ctx.First.Number + ctx.Second.Number
	return &ctx
}

func main() {
	rand.Seed(time.Now().Unix())
	ctx := initCtx()
	http.HandleFunc("/first", ctx.FirstHandler)
	http.HandleFunc("/second", ctx.SecondHandler)
	http.HandleFunc("/summa", ctx.SummaHandler)

	ctx2 := initCtx()
	http.HandleFunc("/f2", ctx2.FirstHandler)
	http.HandleFunc("/s2", ctx2.SecondHandler)
	http.HandleFunc("/summa2", ctx2.SummaHandler)

	http.ListenAndServe(":1234", nil)

}
