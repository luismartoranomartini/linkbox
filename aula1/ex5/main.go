package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Resultado: 42")
	fmt.Fprintln(os.Stderr, "avisos: arquivo ausente")
}
