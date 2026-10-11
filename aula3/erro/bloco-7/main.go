package main

import (
	"errors"
	"fmt"
)

var ErrPilhaVazia = errors.New("pilha vazia")

func main() {
	e := fmt.Errorf("ler topo: %w", ErrPilhaVazia)
	// fmt.Println(e)
	// fmt.Println(e == ErrPilhaVazia)
	// fmt.Println(errors.Unwrap(e) == ErrPilhaVazia)
	// fmt.Printf("%T\n", e)

	// v := fmt.Errorf("ler topo: %v", ErrPilhaVazia)
	// fmt.Println(v)
	// fmt.Println(errors.Unwrap(v))
	// fmt.Printf("%T\n", v)

	e2 := fmt.Errorf("fase 2: %w", e)
	fmt.Println(e2 == ErrPilhaVazia)
	fmt.Println(errors.Is(e, ErrPilhaVazia))
	// fmt.Println(errors.Is(e, ErrOutro))

	fmt.Println("----------------------")
	e3 := fmt.Errorf("fase 2: %w", e)
	fmt.Println(errors.Is(e3, ErrPilhaVazia))
	v := fmt.Errorf("ler topo: %v", ErrPilhaVazia)
	fmt.Println(errors.Is(v, ErrPilhaVazia))

	fmt.Println("----------------------")
	fmt.Println(errors.Is(ErrPilhaVazia, ErrPilhaVazia))
	fmt.Println(errors.Is(nil, ErrPilhaVazia))
	fmt.Println(errors.Is(nil, nil))
}
