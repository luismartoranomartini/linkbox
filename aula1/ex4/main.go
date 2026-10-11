package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
	case "contar":
		cmd := flag.NewFlagSet("contar", flag.ExitOnError)
		ate := cmd.Int("ate", 3, "limite")
		cmd.Parse(os.Args[2:])
		fmt.Println(*ate, cmd.Args())
	}
}
