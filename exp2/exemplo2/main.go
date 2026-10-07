package main

import (
	"fmt"
	"os"
)

func main() {
	buf := make([]byte, 8)
	n, err := os.Stdin.Read(buf)
	fmt.Println(n, string(buf[:n]), err)
	fmt.Fprintln(os.Stdout, "stdout")
	fmt.Fprintln(os.Stderr, "stderr")
}
