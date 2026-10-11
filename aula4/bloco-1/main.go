package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.WriteFile("nota.txt", []byte("café\n"), 0o644)
	fmt.Println("escrever:", err)
	b, err := os.ReadFile("nota.txt")
	fmt.Println("ler:", len(b), err)
	fmt.Printf("%q\n", b)
	_, err = os.ReadFile("nao-existe.txt")
	fmt.Println(err)
}
