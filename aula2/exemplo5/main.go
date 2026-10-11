package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprintln(w, "a")
	fmt.Fprintln(w, "b")
	fmt.Println(w.Buffered())
	w.Flush()
}
