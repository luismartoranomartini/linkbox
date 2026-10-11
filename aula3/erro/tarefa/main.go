package main

import (
	"errors"
	"fmt"
	"os"
)

var ErrUso = errors.New("uso incorreto")

type ErroValidacao struct {
	Campo string
}

func (e *ErroValidacao) Error() string {
	return "campo inválido: " + e.Campo
}

func executar(modo string) error {
	switch modo {
	case "uso":
		return ErrUso
	case "vazio":
		return &ErroValidacao{Campo: "texto"}
	case "es":
		return errors.New("disco cheio")
	}
	return nil
}

func codigoSaida(err error) int {
	if err == nil {
		return 0
	}
	if errors.Is(err, ErrUso) {
		return 2
	}
	var ev *ErroValidacao
	if errors.As(err, &ev) {
		return 3
	}
	return 1
}

func main() {
	err := executar(os.Args[1])
	// err = fmt.Errorf("código: %w", err)
	// fmt.Fprintln(os.Stderr, err)
	if err != nil {
		err = fmt.Errorf("código: %w", err)
		fmt.Fprintln(os.Stderr, err)
	}
	os.Exit(codigoSaida(err))

}
