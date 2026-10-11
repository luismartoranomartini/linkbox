package main

import (
	"fmt"
	"os"
)

func main() {
	os.WriteFile("velho.txt", []byte("antigo"), 0o644)
	os.WriteFile("novo.txt", []byte("recente"), 0o644)
	err := os.Rename("novo.txt", "velho.txt")
	b, _ := os.ReadFile("velho.txt")
	fmt.Println(err, string(b))
	_, err = os.ReadFile("novo.txt")
	fmt.Println(err)
	fmt.Println(os.Rename("nada.txt", "x.txt"))
}
