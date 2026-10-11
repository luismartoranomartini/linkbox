package main

import (
	"encoding/json"
	"fmt"
)

type Item struct {
	Nome string `json:"nome"`
	Qtd  int    `json:"qtd"`
}

var it Item

func main() {
	err := json.Unmarshal([]byte(`{"nome":"caneta","extra":1}`), &it)
	fmt.Println(it, err)
	err = json.Unmarshal([]byte(`{"nome":`), &it)
	fmt.Println(err)
	err = json.Unmarshal([]byte(`{}`), it)
	fmt.Println(err)

	var l []Item
	err = json.Unmarshal([]byte(`[{"nome":"a"},{"nome":"b"}]`), &l)
	fmt.Println(l, err)
	fmt.Println(json.Unmarshal([]byte(""), &l))

}
