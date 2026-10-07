package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// n, err := io.Copy(os.Stdout, strings.NewReader("olá\n"))
	// fmt.Println(n, err)
	//
	// n1, err1 := io.Copy(os.Stdout, strings.NewReader(""))
	// fmt.Println(n1, err1)
	//
	// f, _ := os.Open("/dev/nil")
	// n2, err2 := io.Copy(f, strings.NewReader("abc"))
	// fmt.Println(n2, err2)

	n3, err3 := io.Copy(os.Stdout, os.Stdin)
	fmt.Fprintln(os.Stderr, n3, err3)
}
