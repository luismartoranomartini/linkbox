package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	cru := flag.Bool("cru", false, "copia sem numerar")

	flag.Parse()

	if *cru {
		if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
			fmt.Fprintln(os.Stderr, "numera:", err)
			os.Exit(1)
		}
		return
	}

	// fmt.Fprintln(os.Stderr, "cru:", *cru)

	sc := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	n := 0
	for sc.Scan() {
		n++
		fmt.Fprintf(w, "%3d  %s\n", n, sc.Text())
	}
	if err := w.Flush(); err != nil {
		fmt.Fprintln(os.Stderr, "numera:", err)
		os.Exit(1)
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "numera:", err)
		os.Exit(1)
	}
}
