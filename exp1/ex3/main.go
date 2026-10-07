package main

import (
	"flag"
	"fmt"
)

func main() {
	rotulo := flag.String("rotulo", "item", "prefixo")
	flag.Parse()
	fmt.Println(flag.NArg(), flag.Args())

	for i, a := range flag.Args() {
		fmt.Println(*rotulo, i, a)
	}
}
