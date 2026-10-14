package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("pong\n"))
	// })
	//
	// mux.HandleFunc("POST /ping", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("ping recebido\n"))
	// })

	// mux.HandleFunc("GET /livros/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("livro " + r.PathValue("id") + "\n"))
	// })
	//
	// mux.HandleFunc("GET /arq/{resto...}", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("arq " + r.PathValue("resto") + "\n"))
	// })

	quem := func(w http.ResponseWriter, r *http.Request) { w.Write([]byte(r.Pattern + "\n")) }

	mux.HandleFunc("GET /livros/{id}", quem)
	mux.HandleFunc("GET /livros/novo", quem)
	mux.HandleFunc("/imagens/", quem)
	mux.HandleFunc("/imagens/miniaturas/", quem)
	mux.HandleFunc("/ping", quem)
	mux.HandleFunc("GET /ping", quem)

	server := &http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}
