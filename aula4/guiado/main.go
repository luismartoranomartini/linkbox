package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Nota struct {
	Titulo string `json:"titulo"`
	Peso   int    `json:"peso"`
}

func carregar(arq string) ([]Nota, error) {
	dados, err := os.ReadFile(arq)
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("ler %s: %w", arq, err)
	}
	var notas []Nota
	if err := json.Unmarshal(dados, &notas); err != nil {
		return nil, fmt.Errorf("decodificar %s: %w", arq, err)
	}
	return notas, nil
}

func gravar(arq string, notas []Nota) error {
	dados, err := json.Marshal(notas)
	if err != nil {
		return fmt.Errorf("codificar %s: %w", arq, err)
	}
	return salvar(arq, dados)
}

func salvar(nome string, dados []byte) error {
	tmp := nome + ".temp"
	if err := os.WriteFile(tmp, dados, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, nome)
}

func main() {
	notas, err := carregar("notas.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	notas = append(notas, Nota{"nova", len(notas) + 1})
	if err := gravar("notas.json", notas); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(notas), err)
}
