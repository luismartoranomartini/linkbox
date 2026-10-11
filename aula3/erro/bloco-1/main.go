package main

import (
	"errors"
	"fmt"
)

func main() {
	e1 := errors.New("falhou")
	e2 := errors.New("falhou")
	fmt.Println(e1)
	fmt.Println(e1 == e2)
	fmt.Println(e1 == e1)
}
