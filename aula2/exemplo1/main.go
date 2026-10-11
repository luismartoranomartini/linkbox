package main

import (
	"fmt"
	"strings"
)

func main() {
	r := strings.NewReader("abcdef")
	buf := make([]byte, 4)
	for range 3 {
		n, err := r.Read(buf)
		fmt.Println(n, string(buf[:n]), err)
	}
}
