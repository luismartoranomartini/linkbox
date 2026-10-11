package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	texto := "um dois\n  tres  quatro\n"
	sc := bufio.NewScanner(strings.NewReader(texto))
	for sc.Scan() {
		fmt.Printf("%q\n", sc.Text())
	}
	sc = bufio.NewScanner(strings.NewReader(texto))
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		fmt.Printf("%q\n", sc.Text())
	}
}
