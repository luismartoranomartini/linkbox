package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	sc := bufio.NewScanner(strings.NewReader("um\ndois\n\ntres"))
	for sc.Scan() {
		fmt.Println(len(sc.Text()), sc.Text())
	}
	fmt.Println(sc.Err())
}
