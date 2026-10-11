package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	_, err := os.ReadFile("nao-existe.json")
	fmt.Println(errors.Is(err, os.ErrNotExist))
	_, err = os.ReadFile(".")
	fmt.Println(err, errors.Is(err, os.ErrNotExist))
	_, err = os.ReadFile("nao-existe.json")
	err = fmt.Errorf("ler tarefas: %w", err)
	fmt.Println(errors.Is(err, os.ErrNotExist))
}
