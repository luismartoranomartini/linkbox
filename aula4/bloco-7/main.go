package main

import (
	"fmt"
	"os"
)

func salvar(nome string, dados []byte) error {
	tmp := nome + ".temp"
	if err := os.WriteFile(tmp, dados, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, nome)
}

func main() {
	fmt.Println(salvar("a.json", []byte(`[1]`)))
	fmt.Println(salvar("a.json", []byte(`[1,2]`)))
	b, _ := os.ReadFile("a.json")
	fmt.Println(string(b))
	os.WriteFile("c.json", []byte(`[1]`), 0o644)
	os.Mkdir("c.json.temp", 0o755)
	err := salvar("c.json", []byte(`[2]`))
	b, _ = os.ReadFile("c.json")
	fmt.Println(err, string(b))
}
