package main

import (
	"fmt"
	"io"
)

func main() {
	fmt.Println(io.EOF == io.EOF, io.EOF)
	fmt.Printf("%T\n", io.EOF)
}
