package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(len(os.Args))
	fmt.Println(os.Args[0])

	for i, a := range os.Args[1:] {
		fmt.Println(i, a)
	}
}
