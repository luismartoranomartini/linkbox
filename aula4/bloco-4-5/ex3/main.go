package main

import (
	"encoding/json"
	"fmt"
)

type Ruim struct {
	Nome string `json: "nome"`
}

func main() {
	b, _ := json.Marshal(Ruim{"a"})
	fmt.Println(string(b))
}
