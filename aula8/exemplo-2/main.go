package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func saude(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok\n"))
}

func eco(w http.ResponseWriter, r *http.Request) {
	corpo, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "erro ao ler o corpo da requisição", http.StatusBadRequest)
		return
	}
	w.Write([]byte(corpo))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /saude", saude)

	mux.HandleFunc("POST /eco", eco)

	server := &http.Server{
		Addr:              "127.0.0.1:8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
