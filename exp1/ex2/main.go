package main

import (
	"flag"
	"fmt"
)

func main() {
	vezes := flag.Int("vezes", 1, "quantas vezes repetir")
	texto := flag.String("texto", "ola", "texto a repetir")
	bis := flag.Bool("bis", false, "repetir em dobro")
	flag.Parse()
	if *bis {
		*vezes *= 2
	}
	for range *vezes {
		fmt.Println(*texto)
	}
}
