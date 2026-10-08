package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

const uso = "uso: tarefas <add|list> [opções]"

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
			fmt.Fprintln(os.Stderr, "uso: tarefas add [-prio N] [texto]")
			os.Exit(2)
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
				os.Exit(1)
			}
		}
		if len(texto) == 0 {
			fmt.Fprintln(os.Stderr, "tarefas add: texto vazio")
			os.Exit(2)
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
