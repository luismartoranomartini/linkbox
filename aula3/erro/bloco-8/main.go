package main

import (
	"errors"
	"fmt"
)

type ErroValidacao struct {
	Campo string
}

func (e *ErroValidacao) Error() string {
	return "campo inválido: " + e.Campo
}

func validar(campo string) error {
	if campo == "" {
		return &ErroValidacao{Campo: "nome"} // ou o nome do campo correspondente
	}
	return nil
}

func main() {
	err := fmt.Errorf("add: %w", validar(""))
	fmt.Println(err)
	var ev *ErroValidacao
	fmt.Println(errors.As(err, &ev))
	fmt.Println(ev.Campo)
}
