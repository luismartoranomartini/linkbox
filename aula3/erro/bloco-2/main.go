package main

import (
	"errors"
	"fmt"
)

var ErrPilhaVazia = errors.New("pilha vazia")

func topo(p []int) (int, error) {
	if len(p) == 0 {
		return 0, ErrPilhaVazia
	}
}

func main() {
	var e error
	fmt.Println(e == nil)
	e = errors.New("x")
	fmt.Println(e != nil)
	fmt.Printf("%T\n", e)
}
