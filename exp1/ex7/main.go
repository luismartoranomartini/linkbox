package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "uso: calc <soma|mult>")
		os.Exit(2)
	}
	switch os.Args[1] {
	case "soma":
		cmd := flag.NewFlagSet("soma", flag.ExitOnError)
		a := cmd.Int("a", 0, "primeiro número")
		b := cmd.Int("b", 0, "segundo número")
		cmd.Parse(os.Args[2:])
		if len(cmd.Args()) != 0 {
			fmt.Fprintln(os.Stderr, "uso: calc soma -a N -b N")
			os.Exit(2)
		}
		fmt.Println(*a + *b)
	case "mult":
		cmd := flag.NewFlagSet("mult", flag.ExitOnError)
		a := cmd.Int("a", 1, "primeiro número")
		b := cmd.Int("b", 1, "segundo número")
		cmd.Parse(os.Args[2:])
		if len(cmd.Args()) != 0 {
			fmt.Fprintln(os.Stderr, "uso: calc mult -a N -b N")
			os.Exit(2)
		}
		fmt.Println(*a * *b)
	default:

		fmt.Fprintln(os.Stderr, "subcomando desconhecido:", os.Args[1])
		fmt.Fprintln(os.Stderr, "uso: calc <soma|mult>")
		os.Exit(2)
	}
}
