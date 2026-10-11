package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

const uso = "uso: tarefas <add|list> [opções]"

var ErrUso = errors.New("uso incorreto")

type ErroValidacao struct {
	Campo  string
	Motivo string
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

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, uso)
		os.Exit(2)
	}
	switch os.Args[1] {
	case "add":
		cmd := flag.NewFlagSet("add", flag.ExitOnError)
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

		fmt.Fprintf(os.Stdout, "%s [prio %d]\n", texto, *prio)
	case "list":
		cmd := flag.NewFlagSet("list", flag.ExitOnError)
		bruto := cmd.Bool("bruto", false, "copia o stdin sem numerar")
		cmd.Parse(os.Args[2:])
		if len(cmd.Args()) != 0 {
			fmt.Fprintln(os.Stderr, "uso: tarefas list [-bruto]")
			os.Exit(2)
		}
		if *bruto {
			if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
				fmt.Fprintln(os.Stderr, "tarefas list:", err)
				os.Exit(1)
			}
			return
		}
		sc := bufio.NewScanner(os.Stdin)
		w := bufio.NewWriter(os.Stdout)
		n := 0
		for sc.Scan() {
			n++
			fmt.Fprintf(w, "%2d. %s\n", n, sc.Text())
		}
		if err := w.Flush(); err != nil {
			fmt.Fprintln(os.Stderr, "tarefas list:", err)
			os.Exit(1)
		}
		if err := sc.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "tarefas list:", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "subcomando desconhecido:", os.Args[1])
		fmt.Fprintln(os.Stderr, uso)
		os.Exit(2)
	}
}
