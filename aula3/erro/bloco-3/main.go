package main

import (
	"errors"
	"fmt"
)

var ErrPilhaVazia = errors.New("pilha vazia")

func main() {
	e := fmt.Errorf("ler topo: %w", ErrPilhaVazia)
	fmt.Println(e)
	fmt.Println(e == ErrPilhaVazia)
	fmt.Println(errors.Unwrap(e) == ErrPilhaVazia)
	fmt.Printf("%T\n", e)
}
