package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
)

const uso = "uso: tarefas <add|list> [opções]"

var ErrUso = errors.New("uso incorreto")

type ErroValidacao struct {
	Campo  string
	Motivo string
}

type Tarefa struct {
	Titulo string `json:"titulo"`
	Peso   int    `json:"peso"`
}

func (e *ErroValidacao) Error() string {
	return e.Campo + ": " + e.Motivo
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

func carregar(arq string) ([]Tarefa, error) {
	dados, err := os.ReadFile(arq)
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("ler %s: %w", arq, err)
	}
	var tarefas []Tarefa

	if err := json.Unmarshal(dados, &tarefas); err != nil {
		return nil, fmt.Errorf("decodificar %s: %w", arq, err)
	}
	return tarefas, nil
}

func salvar(nome string, dados []byte) error {
	tmp := nome + ".temp"
	if err := os.WriteFile(tmp, dados, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, nome)
}

func gravar(arq string, tarefas []Tarefa) error {
	dados, err := json.Marshal(tarefas)
	if err != nil {
		return fmt.Errorf("codificar %s: %w", arq, err)
	}
	return salvar(arq, dados)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, uso)
		os.Exit(codigoSaida(ErrUso))
	}
	switch os.Args[1] {
	case "add":
		cmd := flag.NewFlagSet("add", flag.ExitOnError)
		arquivo := cmd.String("arquivo", "tarefas.json", "arquivar dados")
		prio := cmd.Int("prio", 1, "prioridade da tarefas")
		cmd.Parse(os.Args[2:])
		if len(cmd.Args()) > 1 {
			err := fmt.Errorf("%w: esperado [-prio N] [texto]", ErrUso)
			fmt.Fprintln(os.Stderr, "tarefas add:", err)
			os.Exit(codigoSaida(err))
		}

		texto := ""
		if len(cmd.Args()) == 1 {
			texto = cmd.Args()[0]
		} else {
			sc := bufio.NewScanner(os.Stdin)
			if sc.Scan() {
				texto = sc.Text()
			}
			if err := sc.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "tarefas add:", err)
				os.Exit(codigoSaida(err))
			}
		}

		if len(texto) == 0 {
			err := &ErroValidacao{Campo: "texto", Motivo: "não pode ser vazio"}
			fmt.Fprintln(os.Stderr, "tarefas add:", err)
			os.Exit(codigoSaida(err))
		}

		if *prio < 1 || *prio > 3 {
			err := &ErroValidacao{Campo: "prioridade", Motivo: "deve estar entre 1 e 3"}
			fmt.Fprintln(os.Stderr, "tarefas add:", err)
			os.Exit(codigoSaida(err))
		}

		tarefas, err := carregar(*arquivo)
		if err != nil {
			fmt.Fprintln(os.Stderr, "tarefas add:", err)
			os.Exit(codigoSaida(err))
		}
		tarefas = append(tarefas, Tarefa{Titulo: texto, Peso: *prio})
		if err := gravar(*arquivo, tarefas); err != nil {
			fmt.Fprintln(os.Stderr, "tarefas add:", err)
			os.Exit(codigoSaida(err))
		}
		fmt.Fprintf(os.Stdout, "%s [prio %d]\n", texto, *prio)

	case "list":
		cmd := flag.NewFlagSet("list", flag.ExitOnError)
		arquivo := cmd.String("arquivo", "tarefas.json", "arquivar dados")
		cmd.Parse(os.Args[2:])
		if len(cmd.Args()) != 0 {
			fmt.Fprintln(os.Stderr, "uso: tarefas list")
			os.Exit(codigoSaida(ErrUso))
		}
		tarefas, err := carregar(*arquivo)
		if err != nil {
			fmt.Fprintln(os.Stderr, "tarefas list:", err)
			os.Exit(codigoSaida(err))
		}
		w := bufio.NewWriter(os.Stdout)
		n := 0

		for _, t := range tarefas {
			n++
			fmt.Fprintf(w, "%2d. %s [prio %d]\n", n, t.Titulo, t.Peso)
		}

		if err := w.Flush(); err != nil {
			fmt.Fprintln(os.Stderr, "tarefas list:", err)
			os.Exit(codigoSaida(err))
		}

	default:
		fmt.Fprintln(os.Stderr, "subcomando desconhecido:", os.Args[1])
		fmt.Fprintln(os.Stderr, uso)
		os.Exit(codigoSaida(ErrUso))
	}
}
