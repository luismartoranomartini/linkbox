package main

import (
	"errors"
	"fmt"
)

func semSentinela() error {
	return errors.New("pilha vazia")
}

func main() {
	fmt.Println(semSentinela() == semSentinela())
}
