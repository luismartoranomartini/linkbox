package main

import (
	"fmt"
	"io"
	"net/http"
)

func tamanho(w http.ResponseWriter, r *http.Request) {
	corpo, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%d byte\n", len(corpo))
}

func main() {
	http.ListenAndServe("127.0.0.1:8084", http.HandlerFunc(tamanho))
}
