package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	var sb strings.Builder
	w := bufio.NewWriter(&sb)
	fmt.Fprintln(w, "abc")
	fmt.Println(sb.Len(), w.Buffered())
	w.Flush()
	fmt.Println(sb.Len(), w.Buffered())
}
