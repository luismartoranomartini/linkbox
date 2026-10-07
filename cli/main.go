package main

import (
	"flag"
	"fmt"
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
		if len(cmd.Args()) != 1 {
			fmt.Fprintln(os.Stderr, "uso: tarefas add [-prio N] <texto>")
			os.Exit(2)
		}
		fmt.Println("add", *prio, cmd.Arg(0))
	case "list":
		cmd := flag.NewFlagSet("list", flag.ExitOnError)
		todas := cmd.Bool("todas", false, "listar todas as tarefas")
		cmd.Parse(os.Args[2:])
		if len(cmd.Args()) != 0 {
			fmt.Fprintln(os.Stderr, "uso: tarefas list [-todas]")
			os.Exit(2)
		}
		fmt.Println("list", *todas)
	default:
		fmt.Fprintln(os.Stderr, "subcomando desconhecido:", os.Args[1])
		fmt.Fprintln(os.Stderr, uso)
		os.Exit(2)
	}
}
