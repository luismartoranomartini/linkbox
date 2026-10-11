package main

import "os"

func main() {
	d := []byte("x")
	os.WriteFile("a.txt", d, 0o600)
	os.WriteFile("b.txt", d, 0o666)
	os.WriteFile("c.txt", d, 0o755)
	os.WriteFile("a.txt", d, 0o666)
}
