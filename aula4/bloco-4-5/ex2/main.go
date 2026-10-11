package main

import (
	"encoding/json"
	"fmt"
)

type Item struct {
	Nome  string `json:"nome"`
	Qtd   int    `json:"qdt,omitempty"`
	Senha string `json:"-"`
}

func main() {
	b, _ := json.Marshal([]Item{{"caneta", 0, "x"}, {"lápis", 3, "x"}})
	fmt.Println(string(b))
}
