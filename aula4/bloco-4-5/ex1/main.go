package main

import (
	"encoding/json"
	"fmt"
)

type P struct {
	X, Y   int
	Rotulo string
}

var nulo []P

func main() {
	b, err := json.Marshal(P{1, 2, "oculto"})
	fmt.Println(string(b), err)
	b, _ = json.Marshal([]P{{3, 4, ""}, {5, 6, ""}})
	fmt.Println(string(b))
	b, _ = json.Marshal(nulo)
	fmt.Println(string(b))
	b, _ = json.Marshal([]P{})
	fmt.Println(string(b))
	b, _ = json.Marshal("café <ok>")
	fmt.Println(string(b))
	fmt.Println(json.Marshal(make(chan int)))
}
