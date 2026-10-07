package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("defer executado")

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "uso: sai <codigo>")
		os.Exit(3)
	}

	if os.Args[1] == "ok" {
		fmt.Println("fim normal")
		return
	}
	fmt.Println("saindo com 3")
	os.Exit(3)

}
